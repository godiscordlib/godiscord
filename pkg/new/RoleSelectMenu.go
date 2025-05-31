package new

import (
	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
)

func RoleSelectMenu() classes.RoleSelectMenu {
	return classes.RoleSelectMenu{
		classes.BaseSelectMenu{
			Type: enums.ComponentType.RoleSelect,
		},
	}
}
