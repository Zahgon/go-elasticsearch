package tasktypeazureopenai

type TaskTypeAzureOpenAI struct {
	Name string
}

var (
	Textembedding = TaskTypeAzureOpenAI{"text_embedding"}

	Completion = TaskTypeAzureOpenAI{"completion"}

	Chatcompletion = TaskTypeAzureOpenAI{"chat_completion"}
)

func (t TaskTypeAzureOpenAI) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeAzureOpenAI) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeAzureOpenAI) String() string { _ = "STUB: not implemented"; return "" }
