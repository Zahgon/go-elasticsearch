package customtasktype

type CustomTaskType struct {
	Name string
}

var (
	Textembedding = CustomTaskType{"text_embedding"}

	Sparseembedding = CustomTaskType{"sparse_embedding"}

	Rerank = CustomTaskType{"rerank"}

	Completion = CustomTaskType{"completion"}
)

func (c CustomTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CustomTaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CustomTaskType) String() string { _ = "STUB: not implemented"; return "" }
