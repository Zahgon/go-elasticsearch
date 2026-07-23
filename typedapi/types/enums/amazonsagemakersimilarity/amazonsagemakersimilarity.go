package amazonsagemakersimilarity

type AmazonSageMakerSimilarity struct {
	Name string
}

var (
	Cosine = AmazonSageMakerSimilarity{"cosine"}

	Dotproduct = AmazonSageMakerSimilarity{"dot_product"}

	L2norm = AmazonSageMakerSimilarity{"l2_norm"}
)

func (a AmazonSageMakerSimilarity) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AmazonSageMakerSimilarity) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AmazonSageMakerSimilarity) String() string { _ = "STUB: not implemented"; return "" }
