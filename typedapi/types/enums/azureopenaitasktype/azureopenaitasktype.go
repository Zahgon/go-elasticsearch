package azureopenaitasktype

type AzureOpenAITaskType struct {
	Name string
}

var (
	Completion = AzureOpenAITaskType{"completion"}

	Chatcompletion = AzureOpenAITaskType{"chat_completion"}

	Textembedding = AzureOpenAITaskType{"text_embedding"}
)

func (a AzureOpenAITaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AzureOpenAITaskType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AzureOpenAITaskType) String() string { _ = "STUB: not implemented"; return "" }
