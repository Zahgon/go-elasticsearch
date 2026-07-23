package types

type SetSecurityUserProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Properties []string `json:"properties,omitempty"`

	Tag *string `json:"tag,omitempty"`
}

func (s *SetSecurityUserProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSetSecurityUserProcessor() *SetSecurityUserProcessor { _ = "STUB: not implemented"; return nil }

type SetSecurityUserProcessorVariant interface {
	SetSecurityUserProcessorCaster() *SetSecurityUserProcessor
}

func (s *SetSecurityUserProcessor) SetSecurityUserProcessorCaster() *SetSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}
