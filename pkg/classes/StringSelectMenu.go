package classes

type StringSelectMenu struct {
	BaseSelectMenu
	Options []StringSelectOption `json:"options"`
}

func (ssm StringSelectMenu) AddOption(option StringSelectOption) StringSelectMenu {
	ssm.Options = append(ssm.Options, option)
	return ssm
}

func (ssm StringSelectMenu) SetCustomID(customID string) StringSelectMenu {
	ssm.CustomID = customID
	return ssm
}
