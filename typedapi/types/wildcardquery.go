package types

type WildcardQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	CaseInsensitive *bool   `json:"case_insensitive,omitempty"`
	QueryName_      *string `json:"_name,omitempty"`

	Rewrite *string `json:"rewrite,omitempty"`

	Value *string `json:"value,omitempty"`

	Wildcard *string `json:"wildcard,omitempty"`
}

func (s *WildcardQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWildcardQuery() *WildcardQuery { _ = "STUB: not implemented"; return nil }

type WildcardQueryVariant interface {
	WildcardQueryCaster() *WildcardQuery
}

func (s *WildcardQuery) WildcardQueryCaster() *WildcardQuery { _ = "STUB: not implemented"; return nil }
