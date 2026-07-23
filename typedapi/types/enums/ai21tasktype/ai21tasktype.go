package ai21tasktype

type Ai21TaskType struct {
	Name string
}

var (
	Completion = Ai21TaskType{"completion"}

	Chatcompletion = Ai21TaskType{"chat_completion"}
)

func (a Ai21TaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Ai21TaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (a Ai21TaskType) String() string { _ = "STUB: not implemented"; return "" }
