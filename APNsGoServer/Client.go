package APNsGoServer

import (
	"net/http"
	"log"
	"golang.org/x/net/http2"
	"io/ioutil"
	"fmt"
	"bytes"
	"encoding/json"
	"os"
)

const(
	devHost = "https://api.development.push.apple.com:443"
	productionHost = "https://api.push.apple.com:443"
)


func NewPayload(notification Notification) APS {
	alert := Alert{
		Title: notification.Title,
		Body: notification.Message,
	}
	return APS{
		Alert: alert,
		Badge: notification.Badge,
		sound: notification.Sound,
		contentAvailable: notification.ContentAvailable,
		category: notification.Category,
	}
}


func NewHeader(authorization string) Header {
	return Header{
		Authorization: authorization,
		Topic: Conf.Topic,
	}
}

type APNSClient struct {
	URL string
	HTTPClient *http.Client
	Logger *log.Logger
}

func NewClient() (*APNSClient, error) {
	var urlString string
	if Conf.Sandbox {
		urlString = devHost
	} else {
		urlString = productionHost
	}
	tr := &http.Transport{}
	err := http2.ConfigureTransport(tr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client := &http.Client{
		Transport: tr,
	}
	logger := log.New(os.Stdout, "ppush ", log.LstdFlags)
	apnsClient := &APNSClient{
		URL: urlString,
		HTTPClient: client,
		Logger: logger,
	}
	return apnsClient, nil
}

func (c *APNSClient) APNsRequest(token string, header Header, payload APS) (*http.Request, error) {
	URL := fmt.Sprintf("%s/3/device/%s", c.URL, token)
	b, err := json.Marshal(payload.Map())
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", URL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header = header.Map()
	return req, nil
}

func InitAPNSClient() error {
	apnsClient, err := NewClient()
	if err != nil {
		return err
	}
	Client = apnsClient
	return nil
}

func Do(notification Notification) {
	for _, token := range notification.Tokens {
		jwtToken, err := CreateJWT()
		if err != nil {
			fmt.Println(err)
			break
		}
		header := NewHeader(jwtToken)
		payload := NewPayload(notification)
		req, err := Client.APNsRequest(token, header, payload)
		if err != nil {
			fmt.Println(err)
			break
		}
		res, err := Client.HTTPClient.Do(req)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(res.Status)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			var r Response
			if err := json.Unmarshal(body, &r); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(r)
			break
		}

	}
}

func SampleRequest() {
	jwtToken, err := CreateJWT()
	if err != nil {
		fmt.Println(err)
		return
	}

	deviceToken := "e322923a9a506f81c06805a017c83c824f51189bbc0bcce655cb4cd69818773a"
	header := map[string][]string{
		"authorization": {fmt.Sprintf("bearer %s", jwtToken)},
		"apns-topic": {"okysoft.pushNotificationSample"},
	}
	payload := []byte(`{"aps": {"alert": {"title": "title", "body": "body"}}}`)

	// develop
	urlString := "https://api.development.push.apple.com:443"
	// production
	//urlString = "https://api.push.apple.com:443"

	URL := fmt.Sprintf("%s/3/device/%s", urlString, deviceToken)
	req, err := http.NewRequest("POST", URL, bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header = header
	if err != nil {
		fmt.Println(err)
		return
	}
	tr := &http.Transport{}
	if err := http2.ConfigureTransport(tr); err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{
		Transport: tr,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Status)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		var r Response
		if err := json.Unmarshal(body, &r); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(r)
	}
}
