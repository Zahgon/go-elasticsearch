package watsonxservicetype

type WatsonxServiceType struct {
	Name string
}

var (
	Watsonxai = WatsonxServiceType{"watsonxai"}
)

func (w WatsonxServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WatsonxServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (w WatsonxServiceType) String() string { _ = "STUB: not implemented"; return "" }
