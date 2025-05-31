package types

type ButtonType int

const (
	ButtonPrimary ButtonType = iota + 1
	ButtonSecondary
	ButtonSuccess
	ButtonDanger
	ButtonLink
	ButtonPremium
)
