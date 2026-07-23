package contextualaiservicetype

type ContextualAIServiceType struct {
	Name string
}

var (
	Contextualai = ContextualAIServiceType{"contextualai"}
)

func (c ContextualAIServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ContextualAIServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c ContextualAIServiceType) String() string { _ = "STUB: not implemented"; return "" }
