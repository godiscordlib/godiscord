package classes

import (
	"godiscord.foo.ng/lib/pkg/types"
)

type SlashCommandData struct {
	ID          string                `json:"id,omitempty"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Type        types.InteractionType `json:"type"`
	// add Options
}
