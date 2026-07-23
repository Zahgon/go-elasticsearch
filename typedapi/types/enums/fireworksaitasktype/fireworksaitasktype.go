package fireworksaitasktype

type FireworksAITaskType struct {
	Name string
}

var (
	Chatcompletion = FireworksAITaskType{"chat_completion"}

	Completion = FireworksAITaskType{"completion"}

	Textembedding = FireworksAITaskType{"text_embedding"}
)

func (f FireworksAITaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FireworksAITaskType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FireworksAITaskType) String() string { _ = "STUB: not implemented"; return "" }
