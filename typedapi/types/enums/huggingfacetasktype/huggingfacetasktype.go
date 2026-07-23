package huggingfacetasktype

type HuggingFaceTaskType struct {
	Name string
}

var (
	Chatcompletion = HuggingFaceTaskType{"chat_completion"}

	Completion = HuggingFaceTaskType{"completion"}

	Rerank = HuggingFaceTaskType{"rerank"}

	Textembedding = HuggingFaceTaskType{"text_embedding"}
)

func (h HuggingFaceTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HuggingFaceTaskType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (h HuggingFaceTaskType) String() string { _ = "STUB: not implemented"; return "" }
