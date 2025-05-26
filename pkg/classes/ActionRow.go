package classes

type ActionRow struct {
	Type       int             `json:"type"`
	ID         string          `json:"id"`
	Components []BaseComponent `json:"components"`
}

func NewActionRow() ActionRow {
	return ActionRow{
		Type:       1,
		Components: make([]BaseComponent, 1, 5),
	}
}

func (ar ActionRow) SetID(ID string) ActionRow {
	ar.ID = ID
	return ar
}

func (ar ActionRow) AddComponent(Component BaseComponent) ActionRow {
	ar.Components = append(ar.Components, Component)
	return ar
}
