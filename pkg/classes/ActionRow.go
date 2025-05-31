package classes

import (
	"godiscord.foo.ng/lib/pkg/enums"
	"godiscord.foo.ng/lib/pkg/types"
)

type ActionRow struct {
	Type       types.ComponentType `json:"type"`
	ID         int32               `json:"id,omitempty"`
	Components []BaseComponent     `json:"components"`
}

func NewActionRow() ActionRow {
	return ActionRow{
		Type:       enums.ComponentType.ActionRow,
		Components: make([]BaseComponent, 0, 5),
	}
}

func (ar ActionRow) SetID(ID int32) ActionRow {
	ar.ID = ID
	return ar
}

func (ar ActionRow) AddComponent(Component BaseComponent) ActionRow {
	ar.Components = append(ar.Components, Component)
	return ar
}
