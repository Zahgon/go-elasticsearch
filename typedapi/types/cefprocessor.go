package types

type CefProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreEmptyValues *bool `json:"ignore_empty_values,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`

	Timezone *string `json:"timezone,omitempty"`
}

func (s *CefProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCefProcessor() *CefProcessor { _ = "STUB: not implemented"; return nil }

type CefProcessorVariant interface {
	CefProcessorCaster() *CefProcessor
}

func (s *CefProcessor) CefProcessorCaster() *CefProcessor { _ = "STUB: not implemented"; return nil }
