package matchtype

type MatchType struct {
	Name string
}

var (
	Simple = MatchType{"simple"}

	Regex = MatchType{"regex"}
)

func (m MatchType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MatchType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m MatchType) String() string { _ = "STUB: not implemented"; return "" }
