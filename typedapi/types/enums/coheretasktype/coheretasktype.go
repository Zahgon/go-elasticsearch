package coheretasktype

type CohereTaskType struct {
	Name string
}

var (
	Completion = CohereTaskType{"completion"}

	Rerank = CohereTaskType{"rerank"}

	Textembedding = CohereTaskType{"text_embedding"}
)

func (c CohereTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CohereTaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CohereTaskType) String() string { _ = "STUB: not implemented"; return "" }
