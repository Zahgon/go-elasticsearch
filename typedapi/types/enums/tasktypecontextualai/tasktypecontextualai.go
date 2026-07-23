package tasktypecontextualai

type TaskTypeContextualAI struct {
	Name string
}

var (
	Rerank = TaskTypeContextualAI{"rerank"}
)

func (t TaskTypeContextualAI) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeContextualAI) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeContextualAI) String() string { _ = "STUB: not implemented"; return "" }
