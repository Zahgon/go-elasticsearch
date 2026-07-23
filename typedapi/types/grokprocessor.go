package types

type GrokProcessor struct {
	Description *string `json:"description,omitempty"`

	EcsCompatibility *string `json:"ecs_compatibility,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	PatternDefinitions map[string]string `json:"pattern_definitions,omitempty"`

	Patterns []string `json:"patterns"`

	Tag *string `json:"tag,omitempty"`

	TraceMatch *bool `json:"trace_match,omitempty"`

	ValidateOnly *bool `json:"validate_only,omitempty"`
}

func (s *GrokProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGrokProcessor() *GrokProcessor { _ = "STUB: not implemented"; return nil }

type GrokProcessorVariant interface {
	GrokProcessorCaster() *GrokProcessor
}

func (s *GrokProcessor) GrokProcessorCaster() *GrokProcessor { _ = "STUB: not implemented"; return nil }
