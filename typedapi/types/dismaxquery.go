package types

type DisMaxQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Queries    []Query `json:"queries"`
	QueryName_ *string `json:"_name,omitempty"`

	TieBreaker *Float64 `json:"tie_breaker,omitempty"`
}

func (s *DisMaxQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDisMaxQuery() *DisMaxQuery { _ = "STUB: not implemented"; return nil }

type DisMaxQueryVariant interface {
	DisMaxQueryCaster() *DisMaxQuery
}

func (s *DisMaxQuery) DisMaxQueryCaster() *DisMaxQuery { _ = "STUB: not implemented"; return nil }
