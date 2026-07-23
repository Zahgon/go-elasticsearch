package types

type InferenceConfig struct {
	Classification *InferenceConfigClassification `json:"classification,omitempty"`

	Regression *InferenceConfigRegression `json:"regression,omitempty"`
}

func NewInferenceConfig() *InferenceConfig { _ = "STUB: not implemented"; return nil }

type InferenceConfigVariant interface {
	InferenceConfigCaster() *InferenceConfig
}

func (s *InferenceConfig) InferenceConfigCaster() *InferenceConfig {
	_ = "STUB: not implemented"
	return nil
}
