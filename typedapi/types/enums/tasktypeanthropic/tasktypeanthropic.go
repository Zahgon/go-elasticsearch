package tasktypeanthropic

type TaskTypeAnthropic struct {
	Name string
}

var (
	Completion = TaskTypeAnthropic{"completion"}

	Chatcompletion = TaskTypeAnthropic{"chat_completion"}
)

func (t TaskTypeAnthropic) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeAnthropic) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeAnthropic) String() string { _ = "STUB: not implemented"; return "" }
