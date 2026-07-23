package types

type LowercaseNormalizer struct {
	Type string `json:"type,omitempty"`
}

func (s LowercaseNormalizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLowercaseNormalizer() *LowercaseNormalizer { _ = "STUB: not implemented"; return nil }

type LowercaseNormalizerVariant interface {
	LowercaseNormalizerCaster() *LowercaseNormalizer
}

func (s *LowercaseNormalizer) LowercaseNormalizerCaster() *LowercaseNormalizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *LowercaseNormalizer) NormalizerCaster() *Normalizer { _ = "STUB: not implemented"; return nil }
