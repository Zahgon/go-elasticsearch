package types

type LowercaseProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *LowercaseProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewLowercaseProcessor() *LowercaseProcessor { _ = "STUB: not implemented"; return nil }

type LowercaseProcessorVariant interface {
	LowercaseProcessorCaster() *LowercaseProcessor
}

func (s *LowercaseProcessor) LowercaseProcessorCaster() *LowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}
