package fireworksaisimilaritytype

type FireworksAISimilarityType struct {
	Name string
}

var (
	Cosine = FireworksAISimilarityType{"cosine"}

	Dotproduct = FireworksAISimilarityType{"dot_product"}

	L2norm = FireworksAISimilarityType{"l2_norm"}
)

func (f FireworksAISimilarityType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FireworksAISimilarityType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FireworksAISimilarityType) String() string { _ = "STUB: not implemented"; return "" }
