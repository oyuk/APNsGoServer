package APNsGoServer

import (
	"net/http"
	"net"
	"github.com/lestrrat/go-server-starter/listener"
	"fmt"
	"io/ioutil"
	"runtime"
	"encoding/json"
)

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/post", ResponseHandler)
}

func RunServer(server *http.Server) {
	var l net.Listener

	listeners, err := listener.ListenAll()
	if err != nil && err != listener.ErrNoListeningTarget {
		fmt.Println(err)
		return
	}
	if len(listeners) > 0 {
		l = listeners[0]
	}
	if l == nil {
		l, err = net.Listen("tcp", ":8080")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	server.Serve(l)
}

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s%s", r.Host, r.RequestURI)
	var err error
	if r.Method != "POST" {
		sendResponse(w,"HTTP Method must be POST",400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendResponse(w,"Failed to read body",400)
	}
	var notifications Notifications
	if err = json.Unmarshal(body, &notifications); err != nil {
		sendResponse(w,"Invalid json", 400)
	} else {
		go enqueueNotification(notifications)
	}
	sendResponse(w, "ok", 200)
}

func enqueueNotification(notifications Notifications) {
	for _, notification := range notifications.Notification {
		for _, token := range notification.Tokens {
			notification2 := notification
			notification2.Tokens = []string{token}
			NotificationQueue <- notification2
		}
	}
}

func StartNotificationWorkers() {
	NotificationQueue = make(chan Notification)
	for i := 0; i <= runtime.NumCPU(); i++ {
		go pushNotificationWorker()
	}
}

func pushNotificationWorker() {
	for {
		notification := <- NotificationQueue
		Do(notification)
	}
}

func sendResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	m := fmt.Sprintf("{\"message\": \"%s\"}", message)
	w.Write([]byte(m))
}
