package types

type ThaiTokenizer struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *ThaiTokenizer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ThaiTokenizer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewThaiTokenizer() *ThaiTokenizer { _ = "STUB: not implemented"; return nil }

type ThaiTokenizerVariant interface {
	ThaiTokenizerCaster() *ThaiTokenizer
}

func (s *ThaiTokenizer) ThaiTokenizerCaster() *ThaiTokenizer { _ = "STUB: not implemented"; return nil }

func (s *ThaiTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
