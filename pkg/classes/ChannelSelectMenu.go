package classes

import (
	"godiscord.foo.ng/lib/pkg/enums"
	"godiscord.foo.ng/lib/pkg/types"
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
