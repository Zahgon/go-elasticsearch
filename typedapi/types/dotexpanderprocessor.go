package types

type DotExpanderProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Override *bool `json:"override,omitempty"`

	Path *string `json:"path,omitempty"`

	Tag *string `json:"tag,omitempty"`
}

func (s *DotExpanderProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDotExpanderProcessor() *DotExpanderProcessor { _ = "STUB: not implemented"; return nil }

type DotExpanderProcessorVariant interface {
	DotExpanderProcessorCaster() *DotExpanderProcessor
}

func (s *DotExpanderProcessor) DotExpanderProcessorCaster() *DotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}
