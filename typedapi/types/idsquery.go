package types

type IdsQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`

	Values []string `json:"values,omitempty"`
}

func (s *IdsQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIdsQuery() *IdsQuery { _ = "STUB: not implemented"; return nil }

type IdsQueryVariant interface {
	IdsQueryCaster() *IdsQuery
}

func (s *IdsQuery) IdsQueryCaster() *IdsQuery { _ = "STUB: not implemented"; return nil }
