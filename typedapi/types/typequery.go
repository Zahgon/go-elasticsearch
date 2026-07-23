package types

type TypeQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
	Value      string   `json:"value"`
}

func (s *TypeQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTypeQuery() *TypeQuery { _ = "STUB: not implemented"; return nil }

type TypeQueryVariant interface {
	TypeQueryCaster() *TypeQuery
}

func (s *TypeQuery) TypeQueryCaster() *TypeQuery { _ = "STUB: not implemented"; return nil }
