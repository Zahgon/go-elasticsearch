package types

type ModelPlotConfig struct {
	AnnotationsEnabled *bool `json:"annotations_enabled,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	Terms *string `json:"terms,omitempty"`
}

func (s *ModelPlotConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewModelPlotConfig() *ModelPlotConfig { _ = "STUB: not implemented"; return nil }

type ModelPlotConfigVariant interface {
	ModelPlotConfigCaster() *ModelPlotConfig
}

func (s *ModelPlotConfig) ModelPlotConfigCaster() *ModelPlotConfig {
	_ = "STUB: not implemented"
	return nil
}
