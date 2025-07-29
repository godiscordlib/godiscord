package classes

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

type Button struct {
	Type     types.ComponentType `json:"type"`
	Style    types.ButtonType    `json:"style"`
	ID       int                 `json:"id,omitempty"`
	CustomID string              `json:"custom_id"`
	Label    string              `json:"label"`
	Emoji    *Emoji              `json:"emoji,omitempty"`
	SKUID    string              `json:"sku_id,omitempty"`
	URL      string              `json:"url"`
	Disabled bool                `json:"disabled"`
}

func (b Button) SetStyle(style types.ButtonType) Button {
	b.Style = style
	return b
}

func (b Button) SetLabel(label string) Button {
	b.Label = label
	return b
}

func (b Button) SetCustomID(customID string) Button {
	b.CustomID = customID
	return b
}

func (b Button) SetEmoji(emoji Emoji) Button {
	b.Emoji = &emoji
	return b
}
func (b Button) SetSKUID(skuid string) Button {
	b.SKUID = skuid
	return b
}
func (b Button) SetURL(url string) Button {
	b.URL = url
	return b
}
func (b Button) SetDisabled(disabled bool) Button {
	b.Disabled = disabled
	return b
}

func (b Button) GetType() types.ComponentType {
	return b.Type
}
