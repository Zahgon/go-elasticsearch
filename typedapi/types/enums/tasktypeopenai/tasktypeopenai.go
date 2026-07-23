package tasktypeopenai

type TaskTypeOpenAI struct {
	Name string
}

var (
	Textembedding = TaskTypeOpenAI{"text_embedding"}

	Chatcompletion = TaskTypeOpenAI{"chat_completion"}

	Completion = TaskTypeOpenAI{"completion"}

	Embedding = TaskTypeOpenAI{"embedding"}
)

func (t TaskTypeOpenAI) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeOpenAI) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeOpenAI) String() string { _ = "STUB: not implemented"; return "" }
