package classes

import (
	"github.com/godiscordlib/godiscord/pkg/enums"
	"github.com/godiscordlib/godiscord/pkg/types"
)

type SlashCommandIntOption struct {
	SlashCommandOption
	Choices      []SlashCommandChoice `json:"choices,omitempty"` // max 25 Choices
	MinValue     int                  `json:"min_value,omitempty"`
	MaxValue     int                  `json:"max_value,omitempty"`
	Autocomplete bool                 `json:"autocomplete"`
}

func (s SlashCommandIntOption) SetMinValue(v int) SlashCommandIntOption {
	s.MinValue = v
	return s
}
func (s SlashCommandIntOption) SetMaxValue(v int) SlashCommandIntOption {
	s.MaxValue = v
	return s
}

func (s SlashCommandIntOption) AddChoice(choice SlashCommandChoice) SlashCommandIntOption {
	s.Choices = append(s.Choices, choice)
	return s
}

func (s SlashCommandIntOption) GetSCOType() types.SlashCommandOptionType {
	return enums.SlashCommandOptionType.Integer
}
