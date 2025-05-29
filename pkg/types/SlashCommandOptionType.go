package types

type SlashCommandOptionType int

const (
	SCOT_SubCommand SlashCommandOptionType = iota + 1
	SCOT_SubCommandGroup
	SCOT_String
	SCOT_Integer
	SCOT_Boolean
	SCOT_User
	SCOT_Channel
	SCOT_Role
	SCOT_Mentionnable
	SCOT_Number
	SCOT_Attachment
)
