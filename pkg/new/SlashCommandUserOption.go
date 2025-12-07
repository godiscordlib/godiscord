package new

import (
	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
)

func SlashCommandUserOption() classes.SlashCommandUserOption {
	return classes.SlashCommandUserOption{
		SlashCommandOptionInt: classes.SlashCommandOption{
			Type: enums.SlashCommandOptionType.User,
		},
	}
}
