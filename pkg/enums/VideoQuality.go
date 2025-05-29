package enums

import (
	"godiscord.foo.ng/lib/pkg/types"
)

var VideoQuality = struct {
	Auto types.VideoQuality
	Full types.VideoQuality
}{
	Auto: types.VQ_Auto,
	Full: types.VQ_Full,
}
