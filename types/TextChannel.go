package types

type TextChannel struct {
	BaseChannel
	LastMessageID *string
	Topic         *string
	NSFW          *bool
	CategoryID    *string
}
