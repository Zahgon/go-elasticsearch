package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fingerprintdigest"
)

type FingerprintProcessor struct {
	Description *string `json:"description,omitempty"`

	Fields []string `json:"fields"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	Method *fingerprintdigest.FingerprintDigest `json:"method,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Salt *string `json:"salt,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *FingerprintProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFingerprintProcessor() *FingerprintProcessor { _ = "STUB: not implemented"; return nil }

type FingerprintProcessorVariant interface {
	FingerprintProcessorCaster() *FingerprintProcessor
}

func (s *FingerprintProcessor) FingerprintProcessorCaster() *FingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}
