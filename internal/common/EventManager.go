package common

import "sync"

type EventManager struct {
	Handlers map[string][]func(args ...interface{})
	Mutex    sync.RWMutex
}

func NewEventManager() *EventManager {
	return &EventManager{
		Handlers: make(map[string][]func(args ...interface{})),
	}
}

func (e *EventManager) On(Event string, Handler func(args ...interface{})) {
	e.Mutex.Lock()
	defer e.Mutex.Unlock()
	e.Handlers[Event] = append(e.Handlers[Event], Handler)
}

func (e *EventManager) Emit(Event string, Args ...interface{}) {
	e.Mutex.Lock()
	defer e.Mutex.Unlock()

	if handlers, found := e.Handlers[Event]; found {
		for _, handler := range handlers {
			go handler(Args...)
		}
	}
}
