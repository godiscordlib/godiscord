package classes

import (
	"godiscord.foo.ng/lib/internal/types"
)

type Button struct {
	Type     int              `json:"type"`
	Style    types.ButtonType `json:"style"`
	CustomID string           `json:"id"`
	Label    string           `json:"label"`
	Emoji    string           `json:"emoji"`
	SKUID    string           `json:"sku_id,omitempty"`
	URL      string           `json:"url"`
	Disabled bool             `json:"disabled"`
}

func (b Button) GetType() int {
	return b.Type
}
