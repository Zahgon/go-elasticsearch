package types

type NeverCondition struct {
}

func NewNeverCondition() *NeverCondition { _ = "STUB: not implemented"; return nil }

type NeverConditionVariant interface {
	NeverConditionCaster() *NeverCondition
}

func (s *NeverCondition) NeverConditionCaster() *NeverCondition {
	_ = "STUB: not implemented"
	return nil
}
