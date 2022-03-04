package of

type TypeEventFunc = func(r *EventRequest) error
type TypeEventCallbackFunc = func(v interface{}) error

type EventRequest struct {
	Receiver Name
	Type     string
	Value    interface{}
	Callback TypeEventCallbackFunc
}
