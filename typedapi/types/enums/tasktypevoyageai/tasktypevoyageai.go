package tasktypevoyageai

type TaskTypeVoyageAI struct {
	Name string
}

var (
	Textembedding = TaskTypeVoyageAI{"text_embedding"}

	Rerank = TaskTypeVoyageAI{"rerank"}
)

func (t TaskTypeVoyageAI) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeVoyageAI) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeVoyageAI) String() string { _ = "STUB: not implemented"; return "" }
