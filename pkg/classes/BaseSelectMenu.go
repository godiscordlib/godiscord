package classes

import "godiscord.foo.ng/lib/pkg/types"

type BaseSelectMenu struct {
	Type        types.ComponentType `json:"type"`
	CustomID    string              `json:"custom_id"`
	Placeholder string              `json:"placeholder,omitempty"`
	MinValues   int                 `json:"min_values,omitempty"`
	MaxValues   int                 `json:"max_values,omitempty"`
	Disabled    bool                `json:"disabled"`
}

func (bm BaseSelectMenu) GetType() types.ComponentType {
	return bm.Type
}
