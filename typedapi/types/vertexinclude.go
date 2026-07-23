package types

type VertexInclude struct {
	Boost *Float64 `json:"boost,omitempty"`
	Term  string   `json:"term"`
}

func (s *VertexInclude) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewVertexInclude() *VertexInclude { _ = "STUB: not implemented"; return nil }

type VertexIncludeVariant interface {
	VertexIncludeCaster() *VertexInclude
}

func (s *VertexInclude) VertexIncludeCaster() *VertexInclude { _ = "STUB: not implemented"; return nil }
