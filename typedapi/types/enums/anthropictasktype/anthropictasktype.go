package anthropictasktype

type AnthropicTaskType struct {
	Name string
}

var (
	Completion = AnthropicTaskType{"completion"}

	Chatcompletion = AnthropicTaskType{"chat_completion"}
)

func (a AnthropicTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AnthropicTaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (a AnthropicTaskType) String() string { _ = "STUB: not implemented"; return "" }
