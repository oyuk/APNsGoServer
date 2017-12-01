package APNsGoServer

import (
	"fmt"
)

type Header struct {
	Id            string
	Expiration    string
	priority      string
	Topic         string
	Authorization string
	collapseId    string
}

func (h Header) Map() map[string][]string {
	header := make(map[string][]string, 6)
	if h.Id != "" {
		header["apns-id"] = []string{h.Id}
	}
	if h.Expiration != "" {
		header["apns-expiration"] = []string{h.Expiration}
	}
	if h.priority != "" {
		header["apns-priority"] = []string{h.priority}
	}
	if h.Topic != "" {
		header["apns-topic"] = []string{h.Topic}
	}
	if h.Authorization != "" {
		header["authorization"] = []string{fmt.Sprintf("bearer %s", h.Authorization)}
	}
	if h.collapseId != "" {
		header["apns-collapse-id"] = []string{h.collapseId}
	}
	return header
}
