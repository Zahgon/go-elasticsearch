package types

type KeyValueProcessor struct {
	Description *string `json:"description,omitempty"`

	ExcludeKeys []string `json:"exclude_keys,omitempty"`

	Field string `json:"field"`

	FieldSplit string `json:"field_split"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	IncludeKeys []string `json:"include_keys,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Prefix *string `json:"prefix,omitempty"`

	StripBrackets *bool `json:"strip_brackets,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`

	TrimKey *string `json:"trim_key,omitempty"`

	TrimValue *string `json:"trim_value,omitempty"`

	ValueSplit string `json:"value_split"`
}

func (s *KeyValueProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewKeyValueProcessor() *KeyValueProcessor { _ = "STUB: not implemented"; return nil }

type KeyValueProcessorVariant interface {
	KeyValueProcessorCaster() *KeyValueProcessor
}

func (s *KeyValueProcessor) KeyValueProcessorCaster() *KeyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}
