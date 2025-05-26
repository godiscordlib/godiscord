package classes

type BaseInteraction struct {
	Type int `json:"type"`
}

type BaseComponent interface {
	GetType() int
}
