package phoneticruletype

type PhoneticRuleType struct {
	Name string
}

var (
	Approx = PhoneticRuleType{"approx"}

	Exact = PhoneticRuleType{"exact"}
)

func (p PhoneticRuleType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PhoneticRuleType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (p PhoneticRuleType) String() string { _ = "STUB: not implemented"; return "" }
