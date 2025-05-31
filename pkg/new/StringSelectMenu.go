package new

import (
	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
)

func StringSelectMenu() classes.StringSelectMenu {
	return classes.StringSelectMenu{
		classes.BaseSelectMenu{
			Type: enums.ComponentType.StringSelect,
		},
		[]classes.StringSelectOption{},
	}
}
