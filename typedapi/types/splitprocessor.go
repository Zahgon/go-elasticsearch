package types

type SplitProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	PreserveTrailing *bool `json:"preserve_trailing,omitempty"`

	Separator string `json:"separator"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *SplitProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSplitProcessor() *SplitProcessor { _ = "STUB: not implemented"; return nil }

type SplitProcessorVariant interface {
	SplitProcessorCaster() *SplitProcessor
}

func (s *SplitProcessor) SplitProcessorCaster() *SplitProcessor {
	_ = "STUB: not implemented"
	return nil
}
