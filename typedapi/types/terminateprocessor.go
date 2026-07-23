package types

type TerminateProcessor struct {
	Description *string `json:"description,omitempty"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`
}

func (s *TerminateProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTerminateProcessor() *TerminateProcessor { _ = "STUB: not implemented"; return nil }

type TerminateProcessorVariant interface {
	TerminateProcessorCaster() *TerminateProcessor
}

func (s *TerminateProcessor) TerminateProcessorCaster() *TerminateProcessor {
	_ = "STUB: not implemented"
	return nil
}
