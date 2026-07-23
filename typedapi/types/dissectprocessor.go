package types

type DissectProcessor struct {
	AppendSeparator *string `json:"append_separator,omitempty"`

	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Pattern string `json:"pattern"`

	Tag *string `json:"tag,omitempty"`
}

func (s *DissectProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDissectProcessor() *DissectProcessor { _ = "STUB: not implemented"; return nil }

type DissectProcessorVariant interface {
	DissectProcessorCaster() *DissectProcessor
}

func (s *DissectProcessor) DissectProcessorCaster() *DissectProcessor {
	_ = "STUB: not implemented"
	return nil
}
