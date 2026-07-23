package tasktypeai21

type TaskTypeAi21 struct {
	Name string
}

var (
	Completion = TaskTypeAi21{"completion"}

	Chatcompletion = TaskTypeAi21{"chat_completion"}
)

func (t TaskTypeAi21) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeAi21) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeAi21) String() string { _ = "STUB: not implemented"; return "" }
