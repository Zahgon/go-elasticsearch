package tasktypedeepseek

type TaskTypeDeepSeek struct {
	Name string
}

var (
	Completion = TaskTypeDeepSeek{"completion"}

	Chatcompletion = TaskTypeDeepSeek{"chat_completion"}
)

func (t TaskTypeDeepSeek) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeDeepSeek) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeDeepSeek) String() string { _ = "STUB: not implemented"; return "" }
