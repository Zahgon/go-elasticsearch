package types

type NonStationary struct {
	PValue Float64 `json:"p_value"`
	RValue Float64 `json:"r_value"`
	Trend  string  `json:"trend"`
}

func (s *NonStationary) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNonStationary() *NonStationary { _ = "STUB: not implemented"; return nil }
