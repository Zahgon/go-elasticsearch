package amazonsagemakerapi

type AmazonSageMakerApi struct {
	Name string
}

var (
	Openai = AmazonSageMakerApi{"openai"}

	Elastic = AmazonSageMakerApi{"elastic"}
)

func (a AmazonSageMakerApi) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AmazonSageMakerApi) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AmazonSageMakerApi) String() string { _ = "STUB: not implemented"; return "" }
