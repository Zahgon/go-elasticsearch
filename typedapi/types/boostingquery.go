package types

type BoostingQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Negative Query `json:"negative"`

	NegativeBoost Float64 `json:"negative_boost"`

	Positive   Query   `json:"positive"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *BoostingQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBoostingQuery() *BoostingQuery { _ = "STUB: not implemented"; return nil }

type BoostingQueryVariant interface {
	BoostingQueryCaster() *BoostingQuery
}

func (s *BoostingQuery) BoostingQueryCaster() *BoostingQuery { _ = "STUB: not implemented"; return nil }
