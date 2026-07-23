package combinedfieldszeroterms

type CombinedFieldsZeroTerms struct {
	Name string
}

var (
	None = CombinedFieldsZeroTerms{"none"}

	All = CombinedFieldsZeroTerms{"all"}
)

func (c CombinedFieldsZeroTerms) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CombinedFieldsZeroTerms) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CombinedFieldsZeroTerms) String() string { _ = "STUB: not implemented"; return "" }
