package termvectoroption

type TermVectorOption struct {
	Name string
}

var (
	No = TermVectorOption{"no"}

	Yes = TermVectorOption{"yes"}

	Withoffsets = TermVectorOption{"with_offsets"}

	Withpositions = TermVectorOption{"with_positions"}

	Withpositionsoffsets = TermVectorOption{"with_positions_offsets"}

	Withpositionsoffsetspayloads = TermVectorOption{"with_positions_offsets_payloads"}

	Withpositionspayloads = TermVectorOption{"with_positions_payloads"}
)

func (t TermVectorOption) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TermVectorOption) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TermVectorOption) String() string { _ = "STUB: not implemented"; return "" }
