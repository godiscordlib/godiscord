package classes

import (
	"github.com/godiscordlib/godiscord/pkg/enums"
	"github.com/godiscordlib/godiscord/pkg/types"
)

type ChannelSelectMenu struct {
	BaseSelectMenu
	ChannelTypes []types.ChannelType `json:"channel_types"`
}

func NewChannelSelectMenu() ChannelSelectMenu {
	return ChannelSelectMenu{
		BaseSelectMenu{
			Type: enums.ComponentType.ChannelSelect,
		},
		[]types.ChannelType{},
	}
}

func (csm ChannelSelectMenu) SetChannelTypes(channeltypes ...types.ChannelType) ChannelSelectMenu {
	csm.ChannelTypes = channeltypes
	return csm
}

func (csm ChannelSelectMenu) SetCustomID(customID string) ChannelSelectMenu {
	csm.CustomID = customID
	return csm
}
