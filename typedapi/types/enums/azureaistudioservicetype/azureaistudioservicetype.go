package azureaistudioservicetype

type AzureAiStudioServiceType struct {
	Name string
}

var (
	Azureaistudio = AzureAiStudioServiceType{"azureaistudio"}
)

func (a AzureAiStudioServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AzureAiStudioServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AzureAiStudioServiceType) String() string { _ = "STUB: not implemented"; return "" }
