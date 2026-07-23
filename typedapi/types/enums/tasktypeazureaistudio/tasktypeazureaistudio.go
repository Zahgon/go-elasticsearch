package tasktypeazureaistudio

type TaskTypeAzureAIStudio struct {
	Name string
}

var (
	Textembedding = TaskTypeAzureAIStudio{"text_embedding"}

	Completion = TaskTypeAzureAIStudio{"completion"}

	Rerank = TaskTypeAzureAIStudio{"rerank"}
)

func (t TaskTypeAzureAIStudio) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeAzureAIStudio) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeAzureAIStudio) String() string { _ = "STUB: not implemented"; return "" }
