package types

type GlobalAggregation struct {
}

func NewGlobalAggregation() *GlobalAggregation { _ = "STUB: not implemented"; return nil }

type GlobalAggregationVariant interface {
	GlobalAggregationCaster() *GlobalAggregation
}

func (s *GlobalAggregation) GlobalAggregationCaster() *GlobalAggregation {
	_ = "STUB: not implemented"
	return nil
}
