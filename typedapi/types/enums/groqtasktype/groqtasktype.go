package groqtasktype

type GroqTaskType struct {
	Name string
}

var (
	Chatcompletion = GroqTaskType{"chat_completion"}
)

func (g GroqTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GroqTaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GroqTaskType) String() string { _ = "STUB: not implemented"; return "" }
