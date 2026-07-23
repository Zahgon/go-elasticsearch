package azureaistudiotasktype

type AzureAiStudioTaskType struct {
	Name string
}

var (
	Completion = AzureAiStudioTaskType{"completion"}

	Rerank = AzureAiStudioTaskType{"rerank"}

	Textembedding = AzureAiStudioTaskType{"text_embedding"}
)

func (a AzureAiStudioTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AzureAiStudioTaskType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AzureAiStudioTaskType) String() string { _ = "STUB: not implemented"; return "" }
