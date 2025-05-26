package types

type InteractionType int

const (
	IT_ChatInput InteractionType = iota + 1
	IT_User
	IT_Message
	IT_PrimaryEntryPoint
)
