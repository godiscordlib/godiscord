package types

type Emoji struct {
	Base
	Name               string
	Animated           bool
	AllowedRoles       []Role
	Author             User
	RequireColons      bool
	IntegrationManaged bool
	Usable             bool
}
