package new

import (
	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
)

func SlashCommand() classes.SlashCommandData {
	return classes.SlashCommandData{
		Type: enums.ApplicationCommandType.ChatInput,
	}
}
