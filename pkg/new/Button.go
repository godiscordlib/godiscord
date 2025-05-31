package new

import (
	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
)

func Button() classes.Button {
	return classes.Button{
		Type: enums.ComponentType.Button,
	}
}
