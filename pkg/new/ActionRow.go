package new

import (
	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
)

func ActionRow() classes.ActionRow {
	return classes.ActionRow{
		Type:       enums.ComponentType.ActionRow,
		Components: make([]classes.BaseComponent, 0, 5),
	}
}
