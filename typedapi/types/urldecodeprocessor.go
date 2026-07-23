package types

type UrlDecodeProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *UrlDecodeProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUrlDecodeProcessor() *UrlDecodeProcessor { _ = "STUB: not implemented"; return nil }

type UrlDecodeProcessorVariant interface {
	UrlDecodeProcessorCaster() *UrlDecodeProcessor
}

func (s *UrlDecodeProcessor) UrlDecodeProcessorCaster() *UrlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}
