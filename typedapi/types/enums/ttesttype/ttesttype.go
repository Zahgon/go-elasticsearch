package ttesttype

type TTestType struct {
	Name string
}

var (
	Paired = TTestType{"paired"}

	Homoscedastic = TTestType{"homoscedastic"}

	Heteroscedastic = TTestType{"heteroscedastic"}
)

func (t TTestType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TTestType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TTestType) String() string { _ = "STUB: not implemented"; return "" }
