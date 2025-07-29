package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var SlashCommandOptionType = struct {
	SubCommand      types.SlashCommandOptionType
	SubCommandGroup types.SlashCommandOptionType
	String          types.SlashCommandOptionType
	Integer         types.SlashCommandOptionType
	Boolean         types.SlashCommandOptionType
	User            types.SlashCommandOptionType
	Channel         types.SlashCommandOptionType
	Role            types.SlashCommandOptionType
	Mentionnable    types.SlashCommandOptionType
	Number          types.SlashCommandOptionType
	Attachment      types.SlashCommandOptionType
}{
	SubCommand:      types.SCOT_SubCommand,
	SubCommandGroup: types.SCOT_SubCommandGroup,
	String:          types.SCOT_String,
	Integer:         types.SCOT_Integer,
	Boolean:         types.SCOT_Boolean,
	User:            types.SCOT_User,
	Channel:         types.SCOT_Channel,
	Role:            types.SCOT_Role,
	Mentionnable:    types.SCOT_Mentionnable,
	Number:          types.SCOT_Number,
	Attachment:      types.SCOT_Attachment,
}
