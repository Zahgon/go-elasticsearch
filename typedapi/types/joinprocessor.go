package types

type JoinProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Separator string `json:"separator"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *JoinProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJoinProcessor() *JoinProcessor { _ = "STUB: not implemented"; return nil }

type JoinProcessorVariant interface {
	JoinProcessorCaster() *JoinProcessor
}

func (s *JoinProcessor) JoinProcessorCaster() *JoinProcessor { _ = "STUB: not implemented"; return nil }
