package types

type InferenceConfigContainer struct {
	Classification *ClassificationInferenceOptions `json:"classification,omitempty"`

	Regression *RegressionInferenceOptions `json:"regression,omitempty"`
}

func NewInferenceConfigContainer() *InferenceConfigContainer { _ = "STUB: not implemented"; return nil }

type InferenceConfigContainerVariant interface {
	InferenceConfigContainerCaster() *InferenceConfigContainer
}

func (s *InferenceConfigContainer) InferenceConfigContainerCaster() *InferenceConfigContainer {
	_ = "STUB: not implemented"
	return nil
}
