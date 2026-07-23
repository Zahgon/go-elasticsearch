package combinedfieldsoperator

type CombinedFieldsOperator struct {
	Name string
}

var (
	Or = CombinedFieldsOperator{"or"}

	And = CombinedFieldsOperator{"and"}
)

func (c CombinedFieldsOperator) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CombinedFieldsOperator) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CombinedFieldsOperator) String() string { _ = "STUB: not implemented"; return "" }
