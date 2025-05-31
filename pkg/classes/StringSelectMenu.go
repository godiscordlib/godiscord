package classes

import "godiscord.foo.ng/lib/pkg/enums"

type StringSelectMenu struct {
	BaseSelectMenu
	Options []StringSelectOption `json:"options"`
}

func NewStringSelectMenu() StringSelectMenu {
	return StringSelectMenu{
		BaseSelectMenu{
			Type: enums.ComponentType.StringSelect,
		},
		[]StringSelectOption{},
	}
}
func (ssm StringSelectMenu) AddOption(option StringSelectOption) StringSelectMenu {
	ssm.Options = append(ssm.Options, option)
	return ssm
}

func (ssm StringSelectMenu) SetCustomID(customID string) StringSelectMenu {
	ssm.CustomID = customID
	return ssm
}
