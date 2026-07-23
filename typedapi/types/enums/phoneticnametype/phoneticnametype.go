package phoneticnametype

type PhoneticNameType struct {
	Name string
}

var (
	Generic = PhoneticNameType{"generic"}

	Ashkenazi = PhoneticNameType{"ashkenazi"}

	Sephardic = PhoneticNameType{"sephardic"}
)

func (p PhoneticNameType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PhoneticNameType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (p PhoneticNameType) String() string { _ = "STUB: not implemented"; return "" }
