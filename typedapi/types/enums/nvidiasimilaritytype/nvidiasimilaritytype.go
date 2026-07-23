package nvidiasimilaritytype

type NvidiaSimilarityType struct {
	Name string
}

var (
	Cosine = NvidiaSimilarityType{"cosine"}

	Dotproduct = NvidiaSimilarityType{"dot_product"}

	L2norm = NvidiaSimilarityType{"l2_norm"}
)

func (n NvidiaSimilarityType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NvidiaSimilarityType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (n NvidiaSimilarityType) String() string { _ = "STUB: not implemented"; return "" }
