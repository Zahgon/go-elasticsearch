package textquerytype

type TextQueryType struct {
	Name string
}

var (
	Bestfields = TextQueryType{"best_fields"}

	Mostfields = TextQueryType{"most_fields"}

	Crossfields = TextQueryType{"cross_fields"}

	Phrase = TextQueryType{"phrase"}

	Phraseprefix = TextQueryType{"phrase_prefix"}

	Boolprefix = TextQueryType{"bool_prefix"}
)

func (t TextQueryType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TextQueryType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TextQueryType) String() string { _ = "STUB: not implemented"; return "" }
