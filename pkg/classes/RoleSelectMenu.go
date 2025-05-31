package classes

type RoleSelectMenu struct {
	BaseSelectMenu
}

func (csm RoleSelectMenu) SetCustomID(customID string) RoleSelectMenu {
	csm.CustomID = customID
	return csm
}
