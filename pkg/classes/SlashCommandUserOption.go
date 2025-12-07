package classes

import (
	"github.com/godiscordlib/godiscord/pkg/enums"
	"github.com/godiscordlib/godiscord/pkg/types"
)

type SlashCommandUserOption struct {
	SlashCommandOptionInt
}

func (scso SlashCommandUserOption) GetSCOType() types.SlashCommandOptionType {
	return enums.SlashCommandOptionType.String
}
