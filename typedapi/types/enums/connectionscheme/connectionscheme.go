package connectionscheme

type ConnectionScheme struct {
	Name string
}

var (
	Http = ConnectionScheme{"http"}

	Https = ConnectionScheme{"https"}
)

func (c ConnectionScheme) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConnectionScheme) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ConnectionScheme) String() string { _ = "STUB: not implemented"; return "" }
