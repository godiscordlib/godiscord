package common

type Emoji struct {
	Base
	Name               string `json:"name"`
	Animated           bool   `json:"animated"`
	AllowedRoles       []Role `json:"roles"`
	Author             User   `json:"user"`
	RequireColons      bool   `json:"require_colon"`
	IntegrationManaged bool   `json:"managed"`
	Usable             bool   `json:"available"`
}
