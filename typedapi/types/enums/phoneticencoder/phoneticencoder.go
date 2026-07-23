package phoneticencoder

type PhoneticEncoder struct {
	Name string
}

var (
	Metaphone = PhoneticEncoder{"metaphone"}

	Doublemetaphone = PhoneticEncoder{"double_metaphone"}

	Soundex = PhoneticEncoder{"soundex"}

	Refinedsoundex = PhoneticEncoder{"refined_soundex"}

	Caverphone1 = PhoneticEncoder{"caverphone1"}

	Caverphone2 = PhoneticEncoder{"caverphone2"}

	Cologne = PhoneticEncoder{"cologne"}

	Nysiis = PhoneticEncoder{"nysiis"}

	Koelnerphonetik = PhoneticEncoder{"koelnerphonetik"}

	Haasephonetik = PhoneticEncoder{"haasephonetik"}

	Beidermorse = PhoneticEncoder{"beider_morse"}

	Daitchmokotoff = PhoneticEncoder{"daitch_mokotoff"}
)

func (p PhoneticEncoder) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PhoneticEncoder) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (p PhoneticEncoder) String() string { _ = "STUB: not implemented"; return "" }
