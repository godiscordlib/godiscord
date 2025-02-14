package types

type Message struct {
	Base
	ChannelID        string
	Author           User
	Content          *string
	Timestamp        string // ISO8601 timestamp
	MentionsEveryone bool
	UsersMentions    []User
	RolesMentions    []string
	// TODO:
	// Add ChannelMentions
	// Add RoleMentions
	// Add Attachments
	// Add Embed
	// Add Reactions
	// Add Components
	Pinned bool
	Type   int
	Flags  int
}
