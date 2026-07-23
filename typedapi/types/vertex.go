package types

type Vertex struct {
	Depth  int64   `json:"depth"`
	Field  string  `json:"field"`
	Term   string  `json:"term"`
	Weight Float64 `json:"weight"`
}

func (s *Vertex) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewVertex() *Vertex { _ = "STUB: not implemented"; return nil }
