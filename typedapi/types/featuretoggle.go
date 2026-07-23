package types

type FeatureToggle struct {
	Enabled bool `json:"enabled"`
}

func (s *FeatureToggle) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFeatureToggle() *FeatureToggle { _ = "STUB: not implemented"; return nil }
