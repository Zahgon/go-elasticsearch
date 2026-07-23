package tasktypecohere

type TaskTypeCohere struct {
	Name string
}

var (
	Textembedding = TaskTypeCohere{"text_embedding"}

	Rerank = TaskTypeCohere{"rerank"}

	Completion = TaskTypeCohere{"completion"}
)

func (t TaskTypeCohere) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeCohere) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeCohere) String() string { _ = "STUB: not implemented"; return "" }
