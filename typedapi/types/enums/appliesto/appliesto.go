package appliesto

type AppliesTo struct {
	Name string
}

var (
	Actual = AppliesTo{"actual"}

	Typical = AppliesTo{"typical"}

	Difffromtypical = AppliesTo{"diff_from_typical"}

	Time = AppliesTo{"time"}
)

func (a AppliesTo) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AppliesTo) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (a AppliesTo) String() string { _ = "STUB: not implemented"; return "" }
