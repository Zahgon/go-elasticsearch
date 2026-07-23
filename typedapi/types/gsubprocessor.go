package types

type GsubProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Pattern string `json:"pattern"`

	Replacement string `json:"replacement"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *GsubProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGsubProcessor() *GsubProcessor { _ = "STUB: not implemented"; return nil }

type GsubProcessorVariant interface {
	GsubProcessorCaster() *GsubProcessor
}

func (s *GsubProcessor) GsubProcessorCaster() *GsubProcessor { _ = "STUB: not implemented"; return nil }
