package types

type PrefixQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	CaseInsensitive *bool   `json:"case_insensitive,omitempty"`
	QueryName_      *string `json:"_name,omitempty"`

	Rewrite *string `json:"rewrite,omitempty"`

	Value string `json:"value"`
}

func (s *PrefixQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPrefixQuery() *PrefixQuery { _ = "STUB: not implemented"; return nil }

type PrefixQueryVariant interface {
	PrefixQueryCaster() *PrefixQuery
}

func (s *PrefixQuery) PrefixQueryCaster() *PrefixQuery { _ = "STUB: not implemented"; return nil }
