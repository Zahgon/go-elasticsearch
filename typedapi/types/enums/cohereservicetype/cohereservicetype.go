package cohereservicetype

type CohereServiceType struct {
	Name string
}

var (
	Cohere = CohereServiceType{"cohere"}
)

func (c CohereServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CohereServiceType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CohereServiceType) String() string { _ = "STUB: not implemented"; return "" }
