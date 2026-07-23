package tasktypeopenshiftai

type TaskTypeOpenShiftAi struct {
	Name string
}

var (
	Textembedding = TaskTypeOpenShiftAi{"text_embedding"}

	Chatcompletion = TaskTypeOpenShiftAi{"chat_completion"}

	Completion = TaskTypeOpenShiftAi{"completion"}

	Rerank = TaskTypeOpenShiftAi{"rerank"}
)

func (t TaskTypeOpenShiftAi) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeOpenShiftAi) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeOpenShiftAi) String() string { _ = "STUB: not implemented"; return "" }
