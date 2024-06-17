package events

type Event interface {
	Notify(string)
	GetID() string
}

type EventManager struct {
	observers []Event
}

func (e *EventManager) Notify(event string) {
	for _, observer := range e.observers {
		observer.Notify(event)
	}
}

func (e *EventManager) AddEvent(event Event) {
	e.observers = append(e.observers, event)
}

func (e *EventManager) RemoveEvent(observerToRemove Event) []Event {
	ObserverLength := len(e.observers)
	for i, observer := range e.observers {
		if observerToRemove.GetID() == observer.GetID() {
			e.observers[ObserverLength-1], e.observers[i] = e.observers[i], e.observers[ObserverLength-1]
			return e.observers[:ObserverLength-1]
		}
	}
	return e.observers
}

var listeners *EventManager

func GetEventManager() *EventManager {
	if listeners == nil {
		listeners = &EventManager{}
	}
	return listeners
}
