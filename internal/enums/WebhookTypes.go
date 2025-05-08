package enums

type WebhookType int

const (
	WT_Incomming WebhookType = iota + 1
	WT_ChannelFollower
	WT_Application
)
