package types

type ShapeQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	IgnoreUnmapped *bool                      `json:"ignore_unmapped,omitempty"`
	QueryName_     *string                    `json:"_name,omitempty"`
	ShapeQuery     map[string]ShapeFieldQuery `json:"-"`
}

func (s *ShapeQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ShapeQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewShapeQuery() *ShapeQuery { _ = "STUB: not implemented"; return nil }

type ShapeQueryVariant interface {
	ShapeQueryCaster() *ShapeQuery
}

func (s *ShapeQuery) ShapeQueryCaster() *ShapeQuery { _ = "STUB: not implemented"; return nil }
