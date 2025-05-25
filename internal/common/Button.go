package common

import "godiscord.foo.ng/lib/internal/enums"

type Button struct {
	Type     int              `json:"type"`
	Style    enums.ButtonType `json:"style"`
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
