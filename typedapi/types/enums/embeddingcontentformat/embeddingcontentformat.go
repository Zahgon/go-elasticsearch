package embeddingcontentformat

type EmbeddingContentFormat struct {
	Name string
}

var (
	Text = EmbeddingContentFormat{"text"}

	Base64 = EmbeddingContentFormat{"base64"}
)

func (e EmbeddingContentFormat) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EmbeddingContentFormat) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e EmbeddingContentFormat) String() string { _ = "STUB: not implemented"; return "" }
