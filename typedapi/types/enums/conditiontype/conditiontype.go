package conditiontype

type ConditionType struct {
	Name string
}

var (
	Always = ConditionType{"always"}

	Never = ConditionType{"never"}

	Script = ConditionType{"script"}

	Compare = ConditionType{"compare"}

	Arraycompare = ConditionType{"array_compare"}
)

func (c ConditionType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConditionType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ConditionType) String() string { _ = "STUB: not implemented"; return "" }
