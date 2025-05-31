package new

import (
	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
)

func ActionRow() classes.ActionRow {
	return classes.ActionRow{
		Type:       enums.ComponentType.ActionRow,
		Components: make([]classes.BaseComponent, 0, 5),
	}
}
