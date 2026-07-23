package amazonsagemakerservicetype

type AmazonSageMakerServiceType struct {
	Name string
}

var (
	Amazonsagemaker = AmazonSageMakerServiceType{"amazon_sagemaker"}
)

func (a AmazonSageMakerServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AmazonSageMakerServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AmazonSageMakerServiceType) String() string { _ = "STUB: not implemented"; return "" }
