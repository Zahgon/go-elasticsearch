package types

type MatchAllQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
}

func (s *MatchAllQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMatchAllQuery() *MatchAllQuery { _ = "STUB: not implemented"; return nil }

type MatchAllQueryVariant interface {
	MatchAllQueryCaster() *MatchAllQuery
}

func (s *MatchAllQuery) MatchAllQueryCaster() *MatchAllQuery { _ = "STUB: not implemented"; return nil }
