package types

type TermsSetQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`

	MinimumShouldMatchField *string `json:"minimum_should_match_field,omitempty"`

	MinimumShouldMatchScript *Script `json:"minimum_should_match_script,omitempty"`
	QueryName_               *string `json:"_name,omitempty"`

	Terms []FieldValue `json:"terms"`
}

func (s *TermsSetQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermsSetQuery() *TermsSetQuery { _ = "STUB: not implemented"; return nil }

type TermsSetQueryVariant interface {
	TermsSetQueryCaster() *TermsSetQuery
}

func (s *TermsSetQuery) TermsSetQueryCaster() *TermsSetQuery { _ = "STUB: not implemented"; return nil }
