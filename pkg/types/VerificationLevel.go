package types

type VerificationLevel int

const (
	VL_None VerificationLevel = iota
	VL_Low
	VL_Medium
	VL_High
	VL_Max
)
