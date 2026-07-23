package scriptsorttype

type ScriptSortType struct {
	Name string
}

var (
	String = ScriptSortType{"string"}

	Number = ScriptSortType{"number"}

	Version = ScriptSortType{"version"}
)

func (s ScriptSortType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ScriptSortType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ScriptSortType) String() string { _ = "STUB: not implemented"; return "" }
