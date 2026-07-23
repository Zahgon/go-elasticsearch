package types

type TrimProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *TrimProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTrimProcessor() *TrimProcessor { _ = "STUB: not implemented"; return nil }

type TrimProcessorVariant interface {
	TrimProcessorCaster() *TrimProcessor
}

func (s *TrimProcessor) TrimProcessorCaster() *TrimProcessor { _ = "STUB: not implemented"; return nil }
