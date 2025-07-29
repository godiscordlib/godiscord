package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var DefaultForumSortOrder = struct {
	LatestActivity types.DefaultForumSortOrder
	CreationDate   types.DefaultForumSortOrder
}{
	LatestActivity: types.DFSO_LatestActivity,
	CreationDate:   types.DFSO_CreationDate,
}
