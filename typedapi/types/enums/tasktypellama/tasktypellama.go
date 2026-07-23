package tasktypellama

type TaskTypeLlama struct {
	Name string
}

var (
	Textembedding = TaskTypeLlama{"text_embedding"}

	Chatcompletion = TaskTypeLlama{"chat_completion"}

	Completion = TaskTypeLlama{"completion"}
)

func (t TaskTypeLlama) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeLlama) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeLlama) String() string { _ = "STUB: not implemented"; return "" }
