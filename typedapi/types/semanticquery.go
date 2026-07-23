package types

type SemanticQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Field string `json:"field"`

	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *SemanticQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSemanticQuery() *SemanticQuery { _ = "STUB: not implemented"; return nil }

type SemanticQueryVariant interface {
	SemanticQueryCaster() *SemanticQuery
}

func (s *SemanticQuery) SemanticQueryCaster() *SemanticQuery { _ = "STUB: not implemented"; return nil }
