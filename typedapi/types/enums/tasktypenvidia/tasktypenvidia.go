package tasktypenvidia

type TaskTypeNvidia struct {
	Name string
}

var (
	Chatcompletion = TaskTypeNvidia{"chat_completion"}

	Completion = TaskTypeNvidia{"completion"}

	Rerank = TaskTypeNvidia{"rerank"}

	Textembedding = TaskTypeNvidia{"text_embedding"}
)

func (t TaskTypeNvidia) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeNvidia) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeNvidia) String() string { _ = "STUB: not implemented"; return "" }
