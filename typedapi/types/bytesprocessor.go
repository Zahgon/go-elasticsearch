package types

type BytesProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *BytesProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBytesProcessor() *BytesProcessor { _ = "STUB: not implemented"; return nil }

type BytesProcessorVariant interface {
	BytesProcessorCaster() *BytesProcessor
}

func (s *BytesProcessor) BytesProcessorCaster() *BytesProcessor {
	_ = "STUB: not implemented"
	return nil
}
