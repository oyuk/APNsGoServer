package APNsGoServer

type Notification struct {
	Tokens   []string `json:"token"`
	Message  string   `json:"message"`
	Title            string       `json:"title,omitempty"`
	Subtitle         string       `json:"subtitle,omitempty"`
	Badge int `json:"badge,omitempty"`
	Sound string `json:"sound,omitempty"`
	Category string `json:"category,omitempty"`
	ContentAvailable int `json:"content_available,omitempty"`
	Expiry int `json:"expiry,omitempty"`
	Priority int `json:"priority,omitempty"`
	CustomPayload []CustomPayload `json:"custom_payload,omitempty"`
}

type CustomPayload struct {
	Key   string `json:"key"`
	Value string `json:"val"`
}

type Notifications struct {
	Notification []Notification `json:"notifications"`
}