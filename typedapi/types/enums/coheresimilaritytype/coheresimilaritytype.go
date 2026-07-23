package coheresimilaritytype

type CohereSimilarityType struct {
	Name string
}

var (
	Cosine = CohereSimilarityType{"cosine"}

	Dotproduct = CohereSimilarityType{"dot_product"}

	L2norm = CohereSimilarityType{"l2_norm"}
)

func (c CohereSimilarityType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CohereSimilarityType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CohereSimilarityType) String() string { _ = "STUB: not implemented"; return "" }
