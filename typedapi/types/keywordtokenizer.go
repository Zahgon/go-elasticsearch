package types

type KeywordTokenizer struct {
	BufferSize *int    `json:"buffer_size,omitempty"`
	Type       string  `json:"type,omitempty"`
	Version    *string `json:"version,omitempty"`
}

func (s *KeywordTokenizer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s KeywordTokenizer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewKeywordTokenizer() *KeywordTokenizer { _ = "STUB: not implemented"; return nil }

type KeywordTokenizerVariant interface {
	KeywordTokenizerCaster() *KeywordTokenizer
}

func (s *KeywordTokenizer) KeywordTokenizerCaster() *KeywordTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *KeywordTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
