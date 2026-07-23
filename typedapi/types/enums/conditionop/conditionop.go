package conditionop

type ConditionOp struct {
	Name string
}

var (
	Noteq = ConditionOp{"not_eq"}

	Eq = ConditionOp{"eq"}

	Lt = ConditionOp{"lt"}

	Gt = ConditionOp{"gt"}

	Lte = ConditionOp{"lte"}

	Gte = ConditionOp{"gte"}
)

func (c ConditionOp) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConditionOp) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ConditionOp) String() string { _ = "STUB: not implemented"; return "" }
