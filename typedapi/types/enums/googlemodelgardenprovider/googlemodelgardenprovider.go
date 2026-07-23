package googlemodelgardenprovider

type GoogleModelGardenProvider struct {
	Name string
}

var (
	Google = GoogleModelGardenProvider{"google"}

	Anthropic = GoogleModelGardenProvider{"anthropic"}

	Meta = GoogleModelGardenProvider{"meta"}

	Huggingface = GoogleModelGardenProvider{"hugging_face"}

	Mistral = GoogleModelGardenProvider{"mistral"}

	Ai21 = GoogleModelGardenProvider{"ai21"}
)

func (g GoogleModelGardenProvider) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GoogleModelGardenProvider) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (g GoogleModelGardenProvider) String() string { _ = "STUB: not implemented"; return "" }
