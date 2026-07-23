package functionscoremode

type FunctionScoreMode struct {
	Name string
}

var (
	Multiply = FunctionScoreMode{"multiply"}

	Sum = FunctionScoreMode{"sum"}

	Avg = FunctionScoreMode{"avg"}

	First = FunctionScoreMode{"first"}

	Max = FunctionScoreMode{"max"}

	Min = FunctionScoreMode{"min"}
)

func (f FunctionScoreMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FunctionScoreMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f FunctionScoreMode) String() string { _ = "STUB: not implemented"; return "" }
