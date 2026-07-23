package elserservicetype

type ElserServiceType struct {
	Name string
}

var (
	Elser = ElserServiceType{"elser"}
)

func (e ElserServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ElserServiceType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e ElserServiceType) String() string { _ = "STUB: not implemented"; return "" }
