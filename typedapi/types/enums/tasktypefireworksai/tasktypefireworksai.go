package tasktypefireworksai

type TaskTypeFireworksAI struct {
	Name string
}

var (
	Chatcompletion = TaskTypeFireworksAI{"chat_completion"}

	Completion = TaskTypeFireworksAI{"completion"}

	Textembedding = TaskTypeFireworksAI{"text_embedding"}
)

func (t TaskTypeFireworksAI) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeFireworksAI) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeFireworksAI) String() string { _ = "STUB: not implemented"; return "" }
