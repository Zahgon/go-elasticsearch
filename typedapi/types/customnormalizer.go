package types

type CustomNormalizer struct {
	CharFilter []string `json:"char_filter,omitempty"`
	Filter     []string `json:"filter,omitempty"`
	Type       string   `json:"type,omitempty"`
}

func (s CustomNormalizer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewCustomNormalizer() *CustomNormalizer { _ = "STUB: not implemented"; return nil }

type CustomNormalizerVariant interface {
	CustomNormalizerCaster() *CustomNormalizer
}

func (s *CustomNormalizer) CustomNormalizerCaster() *CustomNormalizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *CustomNormalizer) NormalizerCaster() *Normalizer { _ = "STUB: not implemented"; return nil }
