package new

import (
	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/enums"
)

func Button() classes.Button {
	return classes.Button{
		Type: enums.ComponentType.Button,
	}
}
