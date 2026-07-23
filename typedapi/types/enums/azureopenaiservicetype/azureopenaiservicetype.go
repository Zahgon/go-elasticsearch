package azureopenaiservicetype

type AzureOpenAIServiceType struct {
	Name string
}

var (
	Azureopenai = AzureOpenAIServiceType{"azureopenai"}
)

func (a AzureOpenAIServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AzureOpenAIServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AzureOpenAIServiceType) String() string { _ = "STUB: not implemented"; return "" }
