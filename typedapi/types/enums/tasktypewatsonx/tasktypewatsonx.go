package tasktypewatsonx

type TaskTypeWatsonx struct {
	Name string
}

var (
	Textembedding = TaskTypeWatsonx{"text_embedding"}

	Chatcompletion = TaskTypeWatsonx{"chat_completion"}

	Completion = TaskTypeWatsonx{"completion"}
)

func (t TaskTypeWatsonx) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeWatsonx) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeWatsonx) String() string { _ = "STUB: not implemented"; return "" }
