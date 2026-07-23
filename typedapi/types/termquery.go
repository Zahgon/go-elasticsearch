package types

type TermQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	CaseInsensitive *bool   `json:"case_insensitive,omitempty"`
	QueryName_      *string `json:"_name,omitempty"`

	Value FieldValue `json:"value"`
}

func (s *TermQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermQuery() *TermQuery { _ = "STUB: not implemented"; return nil }

type TermQueryVariant interface {
	TermQueryCaster() *TermQuery
}

func (s *TermQuery) TermQueryCaster() *TermQuery { _ = "STUB: not implemented"; return nil }
