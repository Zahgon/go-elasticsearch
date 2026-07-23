package openaitasktype

type OpenAITaskType struct {
	Name string
}

var (
	Chatcompletion = OpenAITaskType{"chat_completion"}

	Completion = OpenAITaskType{"completion"}

	Textembedding = OpenAITaskType{"text_embedding"}

	Embedding = OpenAITaskType{"embedding"}
)

func (o OpenAITaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OpenAITaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (o OpenAITaskType) String() string { _ = "STUB: not implemented"; return "" }
