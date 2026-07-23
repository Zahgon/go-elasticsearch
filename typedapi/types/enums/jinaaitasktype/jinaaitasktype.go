package jinaaitasktype

type JinaAITaskType struct {
	Name string
}

var (
	Embedding = JinaAITaskType{"embedding"}

	Rerank = JinaAITaskType{"rerank"}

	Textembedding = JinaAITaskType{"text_embedding"}
)

func (j JinaAITaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JinaAITaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (j JinaAITaskType) String() string { _ = "STUB: not implemented"; return "" }
