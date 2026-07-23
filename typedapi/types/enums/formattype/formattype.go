package formattype

type FormatType struct {
	Name string
}

var (
	Delimited = FormatType{"delimited"}

	Ndjson = FormatType{"ndjson"}

	Semistructuredtext = FormatType{"semi_structured_text"}

	Xml = FormatType{"xml"}
)

func (f FormatType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FormatType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f FormatType) String() string { _ = "STUB: not implemented"; return "" }
