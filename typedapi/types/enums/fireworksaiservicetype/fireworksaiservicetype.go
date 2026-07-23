package fireworksaiservicetype

type FireworksAIServiceType struct {
	Name string
}

var (
	Fireworksai = FireworksAIServiceType{"fireworksai"}
)

func (f FireworksAIServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FireworksAIServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f FireworksAIServiceType) String() string { _ = "STUB: not implemented"; return "" }
