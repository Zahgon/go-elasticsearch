package types

type TermsQuery struct {
	Boost      *float32                   `json:"boost,omitempty"`
	QueryName_ *string                    `json:"_name,omitempty"`
	TermsQuery map[string]TermsQueryField `json:"-"`
}

func (s *TermsQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s TermsQuery) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewTermsQuery() *TermsQuery { _ = "STUB: not implemented"; return nil }

type TermsQueryVariant interface {
	TermsQueryCaster() *TermsQuery
}

func (s *TermsQuery) TermsQueryCaster() *TermsQuery { _ = "STUB: not implemented"; return nil }
