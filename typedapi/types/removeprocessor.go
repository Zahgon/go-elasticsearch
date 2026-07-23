package types

type RemoveProcessor struct {
	Description *string `json:"description,omitempty"`

	Field []string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	Keep []string `json:"keep,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`
}

func (s *RemoveProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRemoveProcessor() *RemoveProcessor { _ = "STUB: not implemented"; return nil }

type RemoveProcessorVariant interface {
	RemoveProcessorCaster() *RemoveProcessor
}

func (s *RemoveProcessor) RemoveProcessorCaster() *RemoveProcessor {
	_ = "STUB: not implemented"
	return nil
}
