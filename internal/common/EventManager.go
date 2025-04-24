package common

import "sync"

type EventManager struct {
	Handlers map[string][]func(args ...any)
	Mutex    sync.RWMutex
}

func NewEventManager() *EventManager {
	return &EventManager{
		Handlers: make(map[string][]func(args ...any)),
	}
}

func (e *EventManager) On(Event string, Handler func(args ...any)) {
	e.Mutex.Lock()
	defer e.Mutex.Unlock()
	e.Handlers[Event] = append(e.Handlers[Event], Handler)
}

func (e *EventManager) Emit(Event string, Args ...any) {
	e.Mutex.Lock()
	defer e.Mutex.Unlock()

	if handlers, found := e.Handlers[Event]; found {
		for _, handler := range handlers {
			go handler(Args...)
		}
	}
}
