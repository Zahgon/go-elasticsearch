package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/useragentproperty"
)

type UserAgentProcessor struct {
	Description *string `json:"description,omitempty"`

	ExtractDeviceType *bool `json:"extract_device_type,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Properties []useragentproperty.UserAgentProperty `json:"properties,omitempty"`

	RegexFile *string `json:"regex_file,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *UserAgentProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUserAgentProcessor() *UserAgentProcessor { _ = "STUB: not implemented"; return nil }

type UserAgentProcessorVariant interface {
	UserAgentProcessorCaster() *UserAgentProcessor
}

func (s *UserAgentProcessor) UserAgentProcessorCaster() *UserAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}
