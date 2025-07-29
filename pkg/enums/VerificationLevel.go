package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var VerificationLevel = struct {
	None   types.VerificationLevel
	Low    types.VerificationLevel
	Medium types.VerificationLevel
	High   types.VerificationLevel
	Max    types.VerificationLevel
}{
	None:   types.VL_None,
	Low:    types.VL_Low,
	Medium: types.VL_Medium,
	High:   types.VL_High,
	Max:    types.VL_Max,
}
