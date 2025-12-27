package new

import (
	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
)

func RoleSelectMenu() classes.RoleSelectMenu {
	return classes.RoleSelectMenu{
		BaseSelectMenu: classes.BaseSelectMenu{
			Type: enums.ComponentType.RoleSelect,
		},
	}
}
