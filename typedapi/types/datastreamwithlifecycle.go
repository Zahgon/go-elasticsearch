package types

type DataStreamWithLifecycle struct {
	Lifecycle *DataStreamLifecycleWithRollover `json:"lifecycle,omitempty"`
	Name      string                           `json:"name"`
}

func (s *DataStreamWithLifecycle) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamWithLifecycle() *DataStreamWithLifecycle { _ = "STUB: not implemented"; return nil }
