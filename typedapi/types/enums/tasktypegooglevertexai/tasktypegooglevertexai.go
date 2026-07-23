package tasktypegooglevertexai

type TaskTypeGoogleVertexAI struct {
	Name string
}

var (
	Chatcompletion = TaskTypeGoogleVertexAI{"chat_completion"}

	Completion = TaskTypeGoogleVertexAI{"completion"}

	Textembedding = TaskTypeGoogleVertexAI{"text_embedding"}

	Rerank = TaskTypeGoogleVertexAI{"rerank"}
)

func (t TaskTypeGoogleVertexAI) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeGoogleVertexAI) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeGoogleVertexAI) String() string { _ = "STUB: not implemented"; return "" }
