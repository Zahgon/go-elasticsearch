package types

type AttachmentProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	IndexedChars *int64 `json:"indexed_chars,omitempty"`

	IndexedCharsField *string `json:"indexed_chars_field,omitempty"`

	MaxFieldBytes ByteSize `json:"max_field_bytes,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Properties []string `json:"properties,omitempty"`

	RemoveBinary *bool `json:"remove_binary,omitempty"`

	ResourceName *string `json:"resource_name,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *AttachmentProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAttachmentProcessor() *AttachmentProcessor { _ = "STUB: not implemented"; return nil }

type AttachmentProcessorVariant interface {
	AttachmentProcessorCaster() *AttachmentProcessor
}

func (s *AttachmentProcessor) AttachmentProcessorCaster() *AttachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}
