package types

type AmazonSageMakerTaskSettings struct {
	CustomAttributes *string `json:"custom_attributes,omitempty"`

	EnableExplanations *string `json:"enable_explanations,omitempty"`

	InferenceId *string `json:"inference_id,omitempty"`

	SessionId *string `json:"session_id,omitempty"`

	TargetVariant *string `json:"target_variant,omitempty"`
}

func (s *AmazonSageMakerTaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAmazonSageMakerTaskSettings() *AmazonSageMakerTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

type AmazonSageMakerTaskSettingsVariant interface {
	AmazonSageMakerTaskSettingsCaster() *AmazonSageMakerTaskSettings
}

func (s *AmazonSageMakerTaskSettings) AmazonSageMakerTaskSettingsCaster() *AmazonSageMakerTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
