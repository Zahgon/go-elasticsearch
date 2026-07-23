package tasktypegoogleaistudio

type TaskTypeGoogleAIStudio struct {
	Name string
}

var (
	Textembedding = TaskTypeGoogleAIStudio{"text_embedding"}

	Completion = TaskTypeGoogleAIStudio{"completion"}
)

func (t TaskTypeGoogleAIStudio) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeGoogleAIStudio) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TaskTypeGoogleAIStudio) String() string { _ = "STUB: not implemented"; return "" }
