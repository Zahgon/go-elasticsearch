package tasktypehuggingface

type TaskTypeHuggingFace struct {
	Name string
}

var (
	Chatcompletion = TaskTypeHuggingFace{"chat_completion"}

	Completion = TaskTypeHuggingFace{"completion"}

	Rerank = TaskTypeHuggingFace{"rerank"}

	Textembedding = TaskTypeHuggingFace{"text_embedding"}
)

func (t TaskTypeHuggingFace) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeHuggingFace) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeHuggingFace) String() string { _ = "STUB: not implemented"; return "" }
