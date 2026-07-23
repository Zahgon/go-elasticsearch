package mistraltasktype

type MistralTaskType struct {
	Name string
}

var (
	Textembedding = MistralTaskType{"text_embedding"}

	Completion = MistralTaskType{"completion"}

	Chatcompletion = MistralTaskType{"chat_completion"}
)

func (m MistralTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MistralTaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (m MistralTaskType) String() string { _ = "STUB: not implemented"; return "" }
