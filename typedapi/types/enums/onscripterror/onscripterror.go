package onscripterror

type OnScriptError struct {
	Name string
}

var (
	Fail = OnScriptError{"fail"}

	Continue = OnScriptError{"continue"}
)

func (o OnScriptError) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OnScriptError) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (o OnScriptError) String() string { _ = "STUB: not implemented"; return "" }
