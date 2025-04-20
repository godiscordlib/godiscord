package common

type PermissionOverwrite struct {
	Base
	Type  int
	Allow int
	Deny  int
}
