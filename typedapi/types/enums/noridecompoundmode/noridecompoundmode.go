package noridecompoundmode

type NoriDecompoundMode struct {
	Name string
}

var (
	Discard = NoriDecompoundMode{"discard"}

	None = NoriDecompoundMode{"none"}

	Mixed = NoriDecompoundMode{"mixed"}
)

func (n NoriDecompoundMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NoriDecompoundMode) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (n NoriDecompoundMode) String() string { _ = "STUB: not implemented"; return "" }
