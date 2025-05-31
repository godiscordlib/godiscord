package enums

import "godiscord.foo.ng/lib/pkg/types"

var ComponentType = struct {
	ActionRow         types.ComponentType
	Button            types.ComponentType
	StringSelect      types.ComponentType
	TextInput         types.ComponentType
	UserSelect        types.ComponentType
	RoleSelect        types.ComponentType
	MentionableSelect types.ComponentType
	ChannelSelect     types.ComponentType
	Section           types.ComponentType
	TextDisplay       types.ComponentType
	Thumbnail         types.ComponentType
	MediaGallery      types.ComponentType
	File              types.ComponentType
	Seperator         types.ComponentType
	Container         types.ComponentType
}{
	types.CT_ActionRow,
	types.CT_Button,
	types.CT_StringSelect,
	types.CT_TextInput,
	types.CT_UserSelect,
	types.CT_RoleSelect,
	types.CT_MentionableSelect,
	types.CT_ChannelSelect,
	types.CT_Section,
	types.CT_TextDisplay,
	types.CT_Thumbnail,
	types.CT_MediaGallery,
	types.CT_File,
	types.CT_Seperator,
	types.CT_Container,
}
