package types

type RenameProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField string `json:"target_field"`
}

func (s *RenameProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRenameProcessor() *RenameProcessor { _ = "STUB: not implemented"; return nil }

type RenameProcessorVariant interface {
	RenameProcessorCaster() *RenameProcessor
}

func (s *RenameProcessor) RenameProcessorCaster() *RenameProcessor {
	_ = "STUB: not implemented"
	return nil
}
