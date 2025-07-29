package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var ActivityType = struct {
	Playing   types.ActivityType
	Streaming types.ActivityType
	Listening types.ActivityType
	Watching  types.ActivityType
	Custom    types.ActivityType
}{
	Playing:   types.AT_Playing,
	Streaming: types.AT_Streaming,
	Listening: types.AT_Listening,
	Watching:  types.AT_Watching,
	Custom:    types.AT_Custom,
}
