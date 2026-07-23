package types

type MatchNoneQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
}

func (s *MatchNoneQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMatchNoneQuery() *MatchNoneQuery { _ = "STUB: not implemented"; return nil }

type MatchNoneQueryVariant interface {
	MatchNoneQueryCaster() *MatchNoneQuery
}

func (s *MatchNoneQuery) MatchNoneQueryCaster() *MatchNoneQuery {
	_ = "STUB: not implemented"
	return nil
}
