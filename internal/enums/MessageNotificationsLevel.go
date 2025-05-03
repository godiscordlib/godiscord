package enums

type MessageNotificationLevel int

const (
	MNL_AllMessages MessageNotificationLevel = iota
	MNL_OnlyMentions
)
