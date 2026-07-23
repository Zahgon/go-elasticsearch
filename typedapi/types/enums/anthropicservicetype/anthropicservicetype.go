package anthropicservicetype

type AnthropicServiceType struct {
	Name string
}

var (
	Anthropic = AnthropicServiceType{"anthropic"}
)

func (a AnthropicServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AnthropicServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AnthropicServiceType) String() string { _ = "STUB: not implemented"; return "" }
