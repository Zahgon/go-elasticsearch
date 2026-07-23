package types

type UppercaseProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *UppercaseProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUppercaseProcessor() *UppercaseProcessor { _ = "STUB: not implemented"; return nil }

type UppercaseProcessorVariant interface {
	UppercaseProcessorCaster() *UppercaseProcessor
}

func (s *UppercaseProcessor) UppercaseProcessorCaster() *UppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}
