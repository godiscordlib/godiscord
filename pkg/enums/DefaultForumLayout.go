package enums

import (
	"github.com/godiscordlib/godiscord/pkg/types"
)

var DefaultForumLayout = struct {
	NotSet      types.DefaultForumLayout
	ListView    types.DefaultForumLayout
	GalleryView types.DefaultForumLayout
}{
	NotSet:      types.DFL_NotSet,
	ListView:    types.DFL_ListView,
	GalleryView: types.DFL_GalleryView,
}
