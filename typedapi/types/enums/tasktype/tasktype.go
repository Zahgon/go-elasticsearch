package tasktype

type TaskType struct {
	Name string
}

var (
	Sparseembedding = TaskType{"sparse_embedding"}

	Textembedding = TaskType{"text_embedding"}

	Rerank = TaskType{"rerank"}

	Completion = TaskType{"completion"}

	Chatcompletion = TaskType{"chat_completion"}

	Embedding = TaskType{"embedding"}
)

func (t TaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskType) String() string { _ = "STUB: not implemented"; return "" }
