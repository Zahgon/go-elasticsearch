package types

type LetterTokenizer struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *LetterTokenizer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s LetterTokenizer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewLetterTokenizer() *LetterTokenizer { _ = "STUB: not implemented"; return nil }

type LetterTokenizerVariant interface {
	LetterTokenizerCaster() *LetterTokenizer
}

func (s *LetterTokenizer) LetterTokenizerCaster() *LetterTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *LetterTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
