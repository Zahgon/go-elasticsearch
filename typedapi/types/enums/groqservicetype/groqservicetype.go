package groqservicetype

type GroqServiceType struct {
	Name string
}

var (
	Groq = GroqServiceType{"groq"}
)

func (g GroqServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GroqServiceType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (g GroqServiceType) String() string { _ = "STUB: not implemented"; return "" }
