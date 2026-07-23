package types

type EventDataStream struct {
	Name string `json:"name"`
}

func (s *EventDataStream) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEventDataStream() *EventDataStream { _ = "STUB: not implemented"; return nil }
