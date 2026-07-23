package types

type DateIndexNameProcessor struct {
	DateFormats []string `json:"date_formats,omitempty"`

	DateRounding string `json:"date_rounding"`

	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IndexNameFormat *string `json:"index_name_format,omitempty"`

	IndexNamePrefix *string `json:"index_name_prefix,omitempty"`

	Locale *string `json:"locale,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	Timezone *string `json:"timezone,omitempty"`
}

func (s *DateIndexNameProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDateIndexNameProcessor() *DateIndexNameProcessor { _ = "STUB: not implemented"; return nil }

type DateIndexNameProcessorVariant interface {
	DateIndexNameProcessorCaster() *DateIndexNameProcessor
}

func (s *DateIndexNameProcessor) DateIndexNameProcessorCaster() *DateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}
