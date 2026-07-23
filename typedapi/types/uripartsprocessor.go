package types

type UriPartsProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	KeepOriginal *bool `json:"keep_original,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	RemoveIfSuccessful *bool `json:"remove_if_successful,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *UriPartsProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUriPartsProcessor() *UriPartsProcessor { _ = "STUB: not implemented"; return nil }

type UriPartsProcessorVariant interface {
	UriPartsProcessorCaster() *UriPartsProcessor
}

func (s *UriPartsProcessor) UriPartsProcessorCaster() *UriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}
