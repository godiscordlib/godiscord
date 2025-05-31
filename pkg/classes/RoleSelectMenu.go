package classes

import (
	"godiscord.foo.ng/lib/pkg/enums"
)

type RoleSelectMenu struct {
	BaseSelectMenu
}

func NewRoleSelectMenu() RoleSelectMenu {
	return RoleSelectMenu{
		BaseSelectMenu{
			Type: enums.ComponentType.RoleSelect,
		},
	}
}

func (csm RoleSelectMenu) SetCustomID(customID string) RoleSelectMenu {
	csm.CustomID = customID
	return csm
}
