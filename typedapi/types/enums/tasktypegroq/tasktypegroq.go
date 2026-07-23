package tasktypegroq

type TaskTypeGroq struct {
	Name string
}

var (
	Chatcompletion = TaskTypeGroq{"chat_completion"}
)

func (t TaskTypeGroq) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeGroq) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeGroq) String() string { _ = "STUB: not implemented"; return "" }
