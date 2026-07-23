package types

type ForeachProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Processor ProcessorContainer `json:"processor"`

	Tag *string `json:"tag,omitempty"`
}

func (s *ForeachProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewForeachProcessor() *ForeachProcessor { _ = "STUB: not implemented"; return nil }

type ForeachProcessorVariant interface {
	ForeachProcessorCaster() *ForeachProcessor
}

func (s *ForeachProcessor) ForeachProcessorCaster() *ForeachProcessor {
	_ = "STUB: not implemented"
	return nil
}
