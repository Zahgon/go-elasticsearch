package tasktypeelser

type TaskTypeELSER struct {
	Name string
}

var (
	Sparseembedding = TaskTypeELSER{"sparse_embedding"}
)

func (t TaskTypeELSER) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TaskTypeELSER) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TaskTypeELSER) String() string { _ = "STUB: not implemented"; return "" }
