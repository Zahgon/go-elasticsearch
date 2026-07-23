package types

type FeatureEnabled struct {
	Enabled bool `json:"enabled"`
}

func (s *FeatureEnabled) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFeatureEnabled() *FeatureEnabled { _ = "STUB: not implemented"; return nil }

type FeatureEnabledVariant interface {
	FeatureEnabledCaster() *FeatureEnabled
}

func (s *FeatureEnabled) FeatureEnabledCaster() *FeatureEnabled {
	_ = "STUB: not implemented"
	return nil
}
