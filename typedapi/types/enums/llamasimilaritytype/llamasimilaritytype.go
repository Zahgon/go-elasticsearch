package llamasimilaritytype

type LlamaSimilarityType struct {
	Name string
}

var (
	Cosine = LlamaSimilarityType{"cosine"}

	Dotproduct = LlamaSimilarityType{"dot_product"}

	L2norm = LlamaSimilarityType{"l2_norm"}
)

func (l LlamaSimilarityType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *LlamaSimilarityType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (l LlamaSimilarityType) String() string { _ = "STUB: not implemented"; return "" }
