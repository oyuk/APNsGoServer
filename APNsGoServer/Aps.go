package APNsGoServer

type APS struct {
	Alert Alert
	Badge int
	sound string
	contentAvailable int
	category string
	threadId string
}

type Alert struct {
	Title string
	Body string
	TitleLockKey string
	TitleLocArgs string
	ActionLockKey string
	LockKey string
	LockArgs string
	LaunchImage string
}


func (a APS) Map() map[string]interface{} {
	aps := make(map[string]interface{}, 6)
	aps["alert"] = a.Alert.Map()
	if a.Badge >= -1 {
		aps["badge"] = a.Badge
	}
	if a.sound != "" {
		aps["sound"] = a.sound
	}
	if a.contentAvailable >= -1 {
		aps["content-available"] = a.contentAvailable
	}
	if a.category != "" {
		aps["category"] = a.category
	}
	return map[string]interface{}{"aps": aps}
}

func (a Alert) Map() map[string]interface{} {
	alert := make(map[string]interface{}, 8)
	if a.Title != "" {
		alert["title"] = a.Title
	}
	if a.Body != "" {
		alert["body"] = a.Body
	}
	if a.TitleLockKey != "" {
		alert["title-loc-key"] = a.TitleLockKey
	}
	if a.TitleLocArgs != "" {
		alert["title-loc-args"] = a.TitleLocArgs
	}
	if a.ActionLockKey != "" {
		alert["action-loc-key"] = a.ActionLockKey
	}
	if a.LockKey != "" {
		alert["loc-key"] = a.LockKey
	}
	if a.LockArgs != "" {
		alert["loc-args"] = a.LockArgs
	}
	if a.LaunchImage != "" {
		alert["launch-image"] = a.LaunchImage
	}
	return alert
}
