package mistralservicetype

type MistralServiceType struct {
	Name string
}

var (
	Mistral = MistralServiceType{"mistral"}
)

func (m MistralServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MistralServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (m MistralServiceType) String() string { _ = "STUB: not implemented"; return "" }
