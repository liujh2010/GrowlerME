package common

// IEventAdaptor defines the interface between user-defined types and order types in the engine.
type IEventAdaptor interface {
	EventToIOrder(e interface{}) IOrder
	IOrderToEvent(o IOrder) interface{}
}

// EventStatus is indicates the status of the corresponding event.
type EventStatus int8

const (
	// EXECUTE means that the corresponding event is in execution state.
	EXECUTE EventStatus = 0
	// COMPLETE means that the event has completed.
	COMPLETE EventStatus = 1
	// ERROR means that the event has error.
	ERROR EventStatus = 2
)

// IEvent defines the method that the events in the matching engine should have.
type IEvent interface {
	Order() IOrder
	Status() EventStatus
	SetStatus(status EventStatus)
}

// E is the default implementation of IEvent
type E struct {
	order  IOrder
	status EventStatus
}
