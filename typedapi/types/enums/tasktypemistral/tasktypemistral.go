package tasktypemistral

type TaskTypeMistral struct {
	Name string
}

var (
	Textembedding = TaskTypeMistral{"text_embedding"}

	Chatcompletion = TaskTypeMistral{"chat_completion"}

	Completion = TaskTypeMistral{"completion"}
)

func (t TaskTypeMistral) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeMistral) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeMistral) String() string { _ = "STUB: not implemented"; return "" }
