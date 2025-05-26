package classes

type BaseSelectMenu struct {
	BaseInteraction
	Components  []BaseInteraction `json:"components"`
	CustomID    string            `json:"custom_id"`
	Placeholder string            `json:"placeholder,omitempty"`
	MinValues   int               `json:"min_values,omitempty"`
	MaxValues   int               `json:"max_values,omitempty"`
	Disabled    bool              `json:"disabled"`
}
