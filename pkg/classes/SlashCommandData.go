package classes

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

type SlashCommandData struct {
	ID                        string
	Name                      string
	Description               string
	Type                      types.ApplicationCommandType
	DefaultMembersPermissions []types.Permission
	Options                   []SlashCommandOptionInt
}

func (s SlashCommandData) SetName(name string) SlashCommandData {
	s.Name = name
	return s
}
func (s SlashCommandData) SetDescription(desc string) SlashCommandData {
	s.Description = desc
	return s
}
func (s SlashCommandData) SetDefaultMembersPermissions(perms ...types.Permission) SlashCommandData {
	s.DefaultMembersPermissions = append(s.DefaultMembersPermissions, perms...)
	return s
}
func (s SlashCommandData) AddOption(option SlashCommandOptionInt) SlashCommandData {
	s.Options = append(s.Options, option)
	return s
}

type SlashCommandOption struct {
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Type        types.SlashCommandOptionType `json:"type"`
	Required    bool                         `json:"required"`
}

func (s SlashCommandOption) GetSCOType() types.SlashCommandOptionType {
	return s.Type
}

func (s SlashCommandOption) SetName(name string) SlashCommandOption {
	s.Name = name
	return s
}
func (s SlashCommandOption) SetDescription(desc string) SlashCommandOption {
	s.Description = desc
	return s
}
func (s SlashCommandOption) SetRequired(required bool) SlashCommandOption {
	s.Required = required
	return s
}

type SlashCommandOptionInt interface {
	GetSCOType() types.SlashCommandOptionType
}
