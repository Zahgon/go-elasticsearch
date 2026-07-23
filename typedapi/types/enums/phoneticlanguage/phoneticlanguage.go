package phoneticlanguage

type PhoneticLanguage struct {
	Name string
}

var (
	Any = PhoneticLanguage{"any"}

	Common = PhoneticLanguage{"common"}

	Cyrillic = PhoneticLanguage{"cyrillic"}

	English = PhoneticLanguage{"english"}

	French = PhoneticLanguage{"french"}

	German = PhoneticLanguage{"german"}

	Hebrew = PhoneticLanguage{"hebrew"}

	Hungarian = PhoneticLanguage{"hungarian"}

	Polish = PhoneticLanguage{"polish"}

	Romanian = PhoneticLanguage{"romanian"}

	Russian = PhoneticLanguage{"russian"}

	Spanish = PhoneticLanguage{"spanish"}
)

func (p PhoneticLanguage) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PhoneticLanguage) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (p PhoneticLanguage) String() string { _ = "STUB: not implemented"; return "" }
