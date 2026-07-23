package types

type FailProcessor struct {
	Description *string `json:"description,omitempty"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	Message string `json:"message"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`
}

func (s *FailProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFailProcessor() *FailProcessor { _ = "STUB: not implemented"; return nil }

type FailProcessorVariant interface {
	FailProcessorCaster() *FailProcessor
}

func (s *FailProcessor) FailProcessorCaster() *FailProcessor { _ = "STUB: not implemented"; return nil }
