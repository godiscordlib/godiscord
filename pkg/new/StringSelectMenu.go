package new

import (
	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
)

func StringSelectMenu() classes.StringSelectMenu {
	return classes.StringSelectMenu{
		classes.BaseSelectMenu{
			Type: enums.ComponentType.StringSelect,
		},
		[]classes.StringSelectOption{},
	}
}
