package types

type ExistsQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Field      string  `json:"field"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *ExistsQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewExistsQuery() *ExistsQuery { _ = "STUB: not implemented"; return nil }

type ExistsQueryVariant interface {
	ExistsQueryCaster() *ExistsQuery
}

func (s *ExistsQuery) ExistsQueryCaster() *ExistsQuery { _ = "STUB: not implemented"; return nil }
