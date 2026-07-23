package googleaiservicetype

type GoogleAiServiceType struct {
	Name string
}

var (
	Googleaistudio = GoogleAiServiceType{"googleaistudio"}
)

func (g GoogleAiServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GoogleAiServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (g GoogleAiServiceType) String() string { _ = "STUB: not implemented"; return "" }
