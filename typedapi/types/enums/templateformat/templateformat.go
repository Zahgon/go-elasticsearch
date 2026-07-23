package templateformat

type TemplateFormat struct {
	Name string
}

var (
	String = TemplateFormat{"string"}

	Json = TemplateFormat{"json"}
)

func (t TemplateFormat) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TemplateFormat) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TemplateFormat) String() string { _ = "STUB: not implemented"; return "" }
