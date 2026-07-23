package voyageaitasktype

type VoyageAITaskType struct {
	Name string
}

var (
	Textembedding = VoyageAITaskType{"text_embedding"}

	Rerank = VoyageAITaskType{"rerank"}
)

func (v VoyageAITaskType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *VoyageAITaskType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (v VoyageAITaskType) String() string { _ = "STUB: not implemented"; return "" }
