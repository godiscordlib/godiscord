package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var VideoQuality = struct {
	Auto types.VideoQuality
	Full types.VideoQuality
}{
	Auto: types.VQ_Auto,
	Full: types.VQ_Full,
}
