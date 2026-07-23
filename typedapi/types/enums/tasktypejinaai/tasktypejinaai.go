package tasktypejinaai

type TaskTypeJinaAi struct {
	Name string
}

var (
	Embedding = TaskTypeJinaAi{"embedding"}

	Textembedding = TaskTypeJinaAi{"text_embedding"}

	Rerank = TaskTypeJinaAi{"rerank"}
)

func (t TaskTypeJinaAi) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeJinaAi) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeJinaAi) String() string { _ = "STUB: not implemented"; return "" }
