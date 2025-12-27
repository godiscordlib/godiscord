package new

import (
	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
)

func SlashCommandIntOption() classes.SlashCommandIntOption {
	return classes.SlashCommandIntOption{
		SlashCommandOption: classes.SlashCommandOption{
			Type: enums.SlashCommandOptionType.Integer,
		},
	}
}
