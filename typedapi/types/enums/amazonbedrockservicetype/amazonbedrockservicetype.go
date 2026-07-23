package amazonbedrockservicetype

type AmazonBedrockServiceType struct {
	Name string
}

var (
	Amazonbedrock = AmazonBedrockServiceType{"amazonbedrock"}
)

func (a AmazonBedrockServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AmazonBedrockServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AmazonBedrockServiceType) String() string { _ = "STUB: not implemented"; return "" }
