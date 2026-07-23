package scoremode

type ScoreMode struct {
	Name string
}

var (
	Avg = ScoreMode{"avg"}

	Max = ScoreMode{"max"}

	Min = ScoreMode{"min"}

	Multiply = ScoreMode{"multiply"}

	Total = ScoreMode{"total"}
)

func (s ScoreMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ScoreMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ScoreMode) String() string { _ = "STUB: not implemented"; return "" }
