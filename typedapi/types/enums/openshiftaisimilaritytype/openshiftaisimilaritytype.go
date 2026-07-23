package openshiftaisimilaritytype

type OpenShiftAiSimilarityType struct {
	Name string
}

var (
	Cosine = OpenShiftAiSimilarityType{"cosine"}

	Dotproduct = OpenShiftAiSimilarityType{"dot_product"}

	L2norm = OpenShiftAiSimilarityType{"l2_norm"}
)

func (o OpenShiftAiSimilarityType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OpenShiftAiSimilarityType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (o OpenShiftAiSimilarityType) String() string { _ = "STUB: not implemented"; return "" }
