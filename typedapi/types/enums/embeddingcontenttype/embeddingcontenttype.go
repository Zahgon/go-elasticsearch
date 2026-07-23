package embeddingcontenttype

type EmbeddingContentType struct {
	Name string
}

var (
	Text = EmbeddingContentType{"text"}

	Image = EmbeddingContentType{"image"}

	Audio = EmbeddingContentType{"audio"}

	Video = EmbeddingContentType{"video"}

	Pdf = EmbeddingContentType{"pdf"}
)

func (e EmbeddingContentType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EmbeddingContentType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e EmbeddingContentType) String() string { _ = "STUB: not implemented"; return "" }
