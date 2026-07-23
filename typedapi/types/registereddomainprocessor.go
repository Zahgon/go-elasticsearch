package types

type RegisteredDomainProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *RegisteredDomainProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRegisteredDomainProcessor() *RegisteredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

type RegisteredDomainProcessorVariant interface {
	RegisteredDomainProcessorCaster() *RegisteredDomainProcessor
}

func (s *RegisteredDomainProcessor) RegisteredDomainProcessorCaster() *RegisteredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}
