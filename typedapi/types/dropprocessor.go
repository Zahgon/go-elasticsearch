package types

type DropProcessor struct {
	Description *string `json:"description,omitempty"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`
}

func (s *DropProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDropProcessor() *DropProcessor { _ = "STUB: not implemented"; return nil }

type DropProcessorVariant interface {
	DropProcessorCaster() *DropProcessor
}

func (s *DropProcessor) DropProcessorCaster() *DropProcessor { _ = "STUB: not implemented"; return nil }
