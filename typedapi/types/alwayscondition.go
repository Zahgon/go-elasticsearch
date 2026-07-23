package types

type AlwaysCondition struct {
}

func NewAlwaysCondition() *AlwaysCondition { _ = "STUB: not implemented"; return nil }

type AlwaysConditionVariant interface {
	AlwaysConditionCaster() *AlwaysCondition
}

func (s *AlwaysCondition) AlwaysConditionCaster() *AlwaysCondition {
	_ = "STUB: not implemented"
	return nil
}
