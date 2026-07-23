package densevectorsimilarity

type DenseVectorSimilarity struct {
	Name string
}

var (
	Cosine = DenseVectorSimilarity{"cosine"}

	Dotproduct = DenseVectorSimilarity{"dot_product"}

	L2norm = DenseVectorSimilarity{"l2_norm"}

	Maxinnerproduct = DenseVectorSimilarity{"max_inner_product"}
)

func (d DenseVectorSimilarity) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DenseVectorSimilarity) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DenseVectorSimilarity) String() string { _ = "STUB: not implemented"; return "" }
