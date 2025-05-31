package classes

type Emoji struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Animated           bool   `json:"animated,omitempty"`
	AllowedRoles       []Role `json:"roles,omitempty"`
	Author             User   `json:"user,omitempty"`
	RequireColons      bool   `json:"require_colon,omitempty"`
	IntegrationManaged bool   `json:"managed,omitempty"`
	Usable             bool   `json:"available,omitempty"`
}
