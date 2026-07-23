package watsonxtasktype

type WatsonxTaskType struct {
	Name string
}

var (
	Textembedding = WatsonxTaskType{"text_embedding"}

	Rerank = WatsonxTaskType{"rerank"}

	Chatcompletion = WatsonxTaskType{"chat_completion"}

	Completion = WatsonxTaskType{"completion"}
)

func (w WatsonxTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WatsonxTaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (w WatsonxTaskType) String() string { _ = "STUB: not implemented"; return "" }
