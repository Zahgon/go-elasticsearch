package types

type BoolQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Filter []Query `json:"filter,omitempty"`

	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`

	Must []Query `json:"must,omitempty"`

	MustNot    []Query `json:"must_not,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`

	Should []Query `json:"should,omitempty"`
}

func (s *BoolQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBoolQuery() *BoolQuery { _ = "STUB: not implemented"; return nil }

type BoolQueryVariant interface {
	BoolQueryCaster() *BoolQuery
}

func (s *BoolQuery) BoolQueryCaster() *BoolQuery { _ = "STUB: not implemented"; return nil }
