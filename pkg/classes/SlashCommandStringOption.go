package classes

import (
	"github.com/godiscordlib/godiscord/pkg/enums"
	"github.com/godiscordlib/godiscord/pkg/types"
)

type SlashCommandStringOption struct {
	SlashCommandOption
	Choices      []SlashCommandChoice `json:"choices,omitempty"` // max 25 Choices
	MinLength    int                  `json:"min_length,omitempty"`
	MaxLength    int                  `json:"max_length,omitempty"`
	Autocomplete bool                 `json:"autocomplete"`
}

func (s SlashCommandStringOption) AddChoice(choice SlashCommandChoice) SlashCommandStringOption {
	s.Choices = append(s.Choices, choice)
	return s
}
func (s SlashCommandStringOption) SetMinLength(length int) SlashCommandStringOption {
	s.MinLength = length
	return s
}
func (s SlashCommandStringOption) SetMaxLength(length int) SlashCommandStringOption {
	s.MaxLength = length
	return s
}
func (s SlashCommandStringOption) SetAutocomplete(enabled bool) SlashCommandStringOption {
	s.Autocomplete = enabled
	return s
}

func (scso SlashCommandStringOption) GetSCOType() types.SlashCommandOptionType {
	return enums.SlashCommandOptionType.String
}
