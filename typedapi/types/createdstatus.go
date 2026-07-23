package types

type CreatedStatus struct {
	Created bool `json:"created"`
}

func (s *CreatedStatus) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCreatedStatus() *CreatedStatus { _ = "STUB: not implemented"; return nil }
