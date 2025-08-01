package new

import (
	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
)

func SlashCommandStringOption() classes.SlashCommandStringOption {
	return classes.SlashCommandStringOption{
		SlashCommandOption: classes.SlashCommandOption{
			Type: enums.SlashCommandOptionType.String,
		},
	}
}
