package childscoremode

type ChildScoreMode struct {
	Name string
}

var (
	None = ChildScoreMode{"none"}

	Avg = ChildScoreMode{"avg"}

	Sum = ChildScoreMode{"sum"}

	Max = ChildScoreMode{"max"}

	Min = ChildScoreMode{"min"}
)

func (c ChildScoreMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ChildScoreMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ChildScoreMode) String() string { _ = "STUB: not implemented"; return "" }
