package elsertasktype

type ElserTaskType struct {
	Name string
}

var (
	Sparseembedding = ElserTaskType{"sparse_embedding"}
)

func (e ElserTaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ElserTaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e ElserTaskType) String() string { _ = "STUB: not implemented"; return "" }
