package mergetype

type MergeType struct {
	Name string
}

var (
	Index = MergeType{"index"}

	Template = MergeType{"template"}
)

func (m MergeType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MergeType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m MergeType) String() string { _ = "STUB: not implemented"; return "" }
