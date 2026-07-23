package types

type ConstantScoreQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Filter     Query   `json:"filter"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *ConstantScoreQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewConstantScoreQuery() *ConstantScoreQuery { _ = "STUB: not implemented"; return nil }

type ConstantScoreQueryVariant interface {
	ConstantScoreQueryCaster() *ConstantScoreQuery
}

func (s *ConstantScoreQuery) ConstantScoreQueryCaster() *ConstantScoreQuery {
	_ = "STUB: not implemented"
	return nil
}
