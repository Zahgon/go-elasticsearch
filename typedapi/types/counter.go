package types

type Counter struct {
	Active int64 `json:"active"`
	Total  int64 `json:"total"`
}

func (s *Counter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCounter() *Counter { _ = "STUB: not implemented"; return nil }
