package tasktypecustom

type TaskTypeCustom struct {
	Name string
}

var (
	Textembedding = TaskTypeCustom{"text_embedding"}

	Sparseembedding = TaskTypeCustom{"sparse_embedding"}

	Rerank = TaskTypeCustom{"rerank"}

	Completion = TaskTypeCustom{"completion"}
)

func (t TaskTypeCustom) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeCustom) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeCustom) String() string { _ = "STUB: not implemented"; return "" }
