package types

type Processor struct {
	Count *int64 `json:"count,omitempty"`

	Current *int64 `json:"current,omitempty"`

	Failed *int64 `json:"failed,omitempty"`

	TimeInMillis *int64 `json:"time_in_millis,omitempty"`
}

func (s *Processor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewProcessor() *Processor { _ = "STUB: not implemented"; return nil }
