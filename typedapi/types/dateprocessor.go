package types

type DateProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	Formats []string `json:"formats"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	Locale *string `json:"locale,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	OutputFormat *string `json:"output_format,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`

	Timezone *string `json:"timezone,omitempty"`
}

func (s *DateProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDateProcessor() *DateProcessor { _ = "STUB: not implemented"; return nil }

type DateProcessorVariant interface {
	DateProcessorCaster() *DateProcessor
}

func (s *DateProcessor) DateProcessorCaster() *DateProcessor { _ = "STUB: not implemented"; return nil }
