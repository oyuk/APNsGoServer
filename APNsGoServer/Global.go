package APNsGoServer

var (
	Conf Config
	Client *APNSClient
	NotificationQueue chan Notification
)

