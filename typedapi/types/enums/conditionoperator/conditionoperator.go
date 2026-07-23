package conditionoperator

type ConditionOperator struct {
	Name string
}

var (
	Gt = ConditionOperator{"gt"}

	Gte = ConditionOperator{"gte"}

	Lt = ConditionOperator{"lt"}

	Lte = ConditionOperator{"lte"}
)

func (c ConditionOperator) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConditionOperator) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ConditionOperator) String() string { _ = "STUB: not implemented"; return "" }
