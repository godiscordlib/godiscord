package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var ButtonType = struct {
	Primary   types.ButtonType
	Secondary types.ButtonType
	Success   types.ButtonType
	Danger    types.ButtonType
	Link      types.ButtonType
	Premium   types.ButtonType
}{
	Primary:   types.ButtonPrimary,
	Secondary: types.ButtonSecondary,
	Success:   types.ButtonSuccess,
	Danger:    types.ButtonDanger,
	Link:      types.ButtonLink,
	Premium:   types.ButtonPremium,
}
