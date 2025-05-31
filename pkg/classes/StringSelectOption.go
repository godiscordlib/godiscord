package classes

type StringSelectOption struct {
	Label       string  `json:"label"`
	Value       string  `json:"value"`
	Description *string `json:"description,omitempty"`
	Emoji       *Emoji  `json:"emoji,omitempty"`
	Default     *bool   `json:"default,omitempty"`
}

func NewStringSelectOption() StringSelectOption {
	return StringSelectOption{}
}

func (sso StringSelectOption) SetLabel(label string) StringSelectOption {
	sso.Label = label
	return sso
}
func (sso StringSelectOption) SetValue(value string) StringSelectOption {
	sso.Value = value
	return sso
}
func (sso StringSelectOption) SetDefault(value bool) StringSelectOption {
	sso.Default = &value
	return sso
}
func (sso StringSelectOption) SetDescription(description string) StringSelectOption {
	sso.Description = &description
	return sso
}
func (sso StringSelectOption) SetEmoji(emoji Emoji) StringSelectOption {
	sso.Emoji = &emoji
	return sso
}
