package types

type Invocations struct {
	Total int64 `json:"total"`
}

func (s *Invocations) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInvocations() *Invocations { _ = "STUB: not implemented"; return nil }
