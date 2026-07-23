package scriptlanguage

type ScriptLanguage struct {
	Name string
}

var (
	Painless = ScriptLanguage{"painless"}

	Expression = ScriptLanguage{"expression"}

	Mustache = ScriptLanguage{"mustache"}

	Java = ScriptLanguage{"java"}
)

func (s ScriptLanguage) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ScriptLanguage) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ScriptLanguage) String() string { _ = "STUB: not implemented"; return "" }
