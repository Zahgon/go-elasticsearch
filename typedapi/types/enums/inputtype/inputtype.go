package inputtype

type InputType struct {
	Name string
}

var (
	Http = InputType{"http"}

	Search = InputType{"search"}

	Simple = InputType{"simple"}
)

func (i InputType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *InputType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i InputType) String() string { _ = "STUB: not implemented"; return "" }
