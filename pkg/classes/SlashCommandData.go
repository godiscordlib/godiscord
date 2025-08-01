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
	// add Options
}
