package types

type PermissionOverwrite struct {
	Base
	Type  int
	Allow int
	Deny  int
}
