package jinaaisimilaritytype

type JinaAISimilarityType struct {
	Name string
}

var (
	Cosine = JinaAISimilarityType{"cosine"}

	Dotproduct = JinaAISimilarityType{"dot_product"}

	L2norm = JinaAISimilarityType{"l2_norm"}
)

func (j JinaAISimilarityType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JinaAISimilarityType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JinaAISimilarityType) String() string { _ = "STUB: not implemented"; return "" }
