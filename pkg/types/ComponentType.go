package types

type ComponentType int

const (
	CT_ActionRow ComponentType = iota + 1
	CT_Button
	CT_StringSelect
	CT_TextInput
	CT_UserSelect
	CT_RoleSelect
	CT_MentionableSelect
	CT_ChannelSelect
	CT_Section
	CT_TextDisplay
	CT_Thumbnail
	CT_MediaGallery
	CT_File
	CT_Seperator
)

const (
	CT_Container ComponentType = 17
)
