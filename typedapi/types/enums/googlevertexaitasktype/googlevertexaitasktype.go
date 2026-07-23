package googlevertexaitasktype

type GoogleVertexAITaskType struct {
	Name string
}

var (
	Rerank = GoogleVertexAITaskType{"rerank"}

	Textembedding = GoogleVertexAITaskType{"text_embedding"}

	Completion = GoogleVertexAITaskType{"completion"}

	Chatcompletion = GoogleVertexAITaskType{"chat_completion"}
)

func (g GoogleVertexAITaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GoogleVertexAITaskType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (g GoogleVertexAITaskType) String() string { _ = "STUB: not implemented"; return "" }
