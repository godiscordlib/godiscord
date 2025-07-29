package classes

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

type SlashCommandData struct {
	ID          string                       `json:"id,omitempty"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Type        types.ApplicationCommandType `json:"type"`
	// add Options
}
