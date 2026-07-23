package voyageaiservicetype

type VoyageAIServiceType struct {
	Name string
}

var (
	Voyageai = VoyageAIServiceType{"voyageai"}
)

func (v VoyageAIServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *VoyageAIServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (v VoyageAIServiceType) String() string { _ = "STUB: not implemented"; return "" }
