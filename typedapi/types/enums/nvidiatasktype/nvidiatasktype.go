package nvidiatasktype

type NvidiaTaskType struct {
	Name string
}

var (
	Chatcompletion = NvidiaTaskType{"chat_completion"}

	Completion = NvidiaTaskType{"completion"}

	Rerank = NvidiaTaskType{"rerank"}

	Textembedding = NvidiaTaskType{"text_embedding"}
)

func (n NvidiaTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NvidiaTaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (n NvidiaTaskType) String() string { _ = "STUB: not implemented"; return "" }
