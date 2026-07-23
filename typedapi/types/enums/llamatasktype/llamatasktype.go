package llamatasktype

type LlamaTaskType struct {
	Name string
}

var (
	Textembedding = LlamaTaskType{"text_embedding"}

	Completion = LlamaTaskType{"completion"}

	Chatcompletion = LlamaTaskType{"chat_completion"}
)

func (l LlamaTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *LlamaTaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (l LlamaTaskType) String() string { _ = "STUB: not implemented"; return "" }
