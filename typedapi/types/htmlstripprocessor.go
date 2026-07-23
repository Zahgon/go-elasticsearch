package types

type HtmlStripProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *HtmlStripProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHtmlStripProcessor() *HtmlStripProcessor { _ = "STUB: not implemented"; return nil }

type HtmlStripProcessorVariant interface {
	HtmlStripProcessorCaster() *HtmlStripProcessor
}

func (s *HtmlStripProcessor) HtmlStripProcessorCaster() *HtmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}
