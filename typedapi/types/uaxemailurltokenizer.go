package types

type UaxEmailUrlTokenizer struct {
	MaxTokenLength *int    `json:"max_token_length,omitempty"`
	Type           string  `json:"type,omitempty"`
	Version        *string `json:"version,omitempty"`
}

func (s *UaxEmailUrlTokenizer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s UaxEmailUrlTokenizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewUaxEmailUrlTokenizer() *UaxEmailUrlTokenizer { _ = "STUB: not implemented"; return nil }

type UaxEmailUrlTokenizerVariant interface {
	UaxEmailUrlTokenizerCaster() *UaxEmailUrlTokenizer
}

func (s *UaxEmailUrlTokenizer) UaxEmailUrlTokenizerCaster() *UaxEmailUrlTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *UaxEmailUrlTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
