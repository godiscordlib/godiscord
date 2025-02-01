package types

import "godiscord/enums"

type Role struct {
	Base
	Name               string
	Color              int
	Hoisted            bool
	IconHash           *string
	Position           int
	Permissions        []enums.Permission
	IntegrationManaged bool
	Mentionable        bool
	Flags              int
}
