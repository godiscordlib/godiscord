package classes

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

type ActionRow struct {
	Type       types.ComponentType `json:"type"`
	ID         int32               `json:"id,omitempty"`
	Components []BaseComponent     `json:"components"`
}

func (ar ActionRow) SetID(ID int32) ActionRow {
	ar.ID = ID
	return ar
}

func (ar ActionRow) AddComponent(Component BaseComponent) ActionRow {
	ar.Components = append(ar.Components, Component)
	return ar
}
