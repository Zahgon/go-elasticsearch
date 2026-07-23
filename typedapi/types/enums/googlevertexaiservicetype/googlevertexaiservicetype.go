package googlevertexaiservicetype

type GoogleVertexAIServiceType struct {
	Name string
}

var (
	Googlevertexai = GoogleVertexAIServiceType{"googlevertexai"}
)

func (g GoogleVertexAIServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GoogleVertexAIServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (g GoogleVertexAIServiceType) String() string { _ = "STUB: not implemented"; return "" }
