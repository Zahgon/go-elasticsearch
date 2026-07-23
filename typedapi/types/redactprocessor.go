package types

type RedactProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure          []ProcessorContainer `json:"on_failure,omitempty"`
	PatternDefinitions map[string]string    `json:"pattern_definitions,omitempty"`

	Patterns []string `json:"patterns"`

	Prefix *string `json:"prefix,omitempty"`

	SkipIfUnlicensed *bool `json:"skip_if_unlicensed,omitempty"`

	Suffix *string `json:"suffix,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TraceRedact *bool `json:"trace_redact,omitempty"`
}

func (s *RedactProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRedactProcessor() *RedactProcessor { _ = "STUB: not implemented"; return nil }

type RedactProcessorVariant interface {
	RedactProcessorCaster() *RedactProcessor
}

func (s *RedactProcessor) RedactProcessorCaster() *RedactProcessor {
	_ = "STUB: not implemented"
	return nil
}
