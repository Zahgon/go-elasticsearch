package openshiftaitasktype

type OpenShiftAiTaskType struct {
	Name string
}

var (
	Textembedding = OpenShiftAiTaskType{"text_embedding"}

	Completion = OpenShiftAiTaskType{"completion"}

	Chatcompletion = OpenShiftAiTaskType{"chat_completion"}

	Rerank = OpenShiftAiTaskType{"rerank"}
)

func (o OpenShiftAiTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OpenShiftAiTaskType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (o OpenShiftAiTaskType) String() string { _ = "STUB: not implemented"; return "" }
