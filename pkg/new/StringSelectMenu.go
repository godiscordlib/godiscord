package new

import (
	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
)

func StringSelectMenu() classes.StringSelectMenu {
	return classes.StringSelectMenu{
		BaseSelectMenu: classes.BaseSelectMenu{
			Type: enums.ComponentType.StringSelect,
		},
		Options: []classes.StringSelectOption{},
	}
}
