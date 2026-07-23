package displaytype

type DisplayType struct {
	Name string
}

var (
	Textbox = DisplayType{"textbox"}

	Textarea = DisplayType{"textarea"}

	Numeric = DisplayType{"numeric"}

	Toggle = DisplayType{"toggle"}

	Dropdown = DisplayType{"dropdown"}
)

func (d DisplayType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DisplayType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (d DisplayType) String() string { _ = "STUB: not implemented"; return "" }
