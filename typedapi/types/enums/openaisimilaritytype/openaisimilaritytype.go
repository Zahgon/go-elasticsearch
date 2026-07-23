package openaisimilaritytype

type OpenAISimilarityType struct {
	Name string
}

var (
	Cosine = OpenAISimilarityType{"cosine"}

	Dotproduct = OpenAISimilarityType{"dot_product"}

	L2norm = OpenAISimilarityType{"l2_norm"}
)

func (o OpenAISimilarityType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OpenAISimilarityType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (o OpenAISimilarityType) String() string { _ = "STUB: not implemented"; return "" }
