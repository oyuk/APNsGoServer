package APNsGoServer

import (
	"fmt"
	"net/http"
)

func StartAPNsGoServer(confPath *string) {
	conf, err := LoadConfig(confPath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Loading conf file is error.")
	}
	Conf = conf
	if err != nil {
		panic(err)
	}

	if err := InitAPNSClient(); err != nil {
		panic(err)
	}

	StartNotificationWorkers()
	startServer()
}

func startServer() {
	mux := http.NewServeMux()
	RegisterHandlers(mux)

	server := &http.Server{
		Handler: mux,
	}
	RunServer(server)
}