package openaiservicetype

type OpenAIServiceType struct {
	Name string
}

var (
	Openai = OpenAIServiceType{"openai"}
)

func (o OpenAIServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OpenAIServiceType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (o OpenAIServiceType) String() string { _ = "STUB: not implemented"; return "" }
