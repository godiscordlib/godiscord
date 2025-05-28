package types

type ApplicationCommandType int

const (
	ACT_ChatInput ApplicationCommandType = iota + 1
	ACT_User
	ACT_Message
	ACT_PrimaryEntryPoint
)
