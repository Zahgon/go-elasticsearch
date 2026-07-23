package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jsonprocessorconflictstrategy"
)

type JsonProcessor struct {
	AddToRoot *bool `json:"add_to_root,omitempty"`

	AddToRootConflictStrategy *jsonprocessorconflictstrategy.JsonProcessorConflictStrategy `json:"add_to_root_conflict_strategy,omitempty"`

	AllowDuplicateKeys *bool `json:"allow_duplicate_keys,omitempty"`

	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *JsonProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJsonProcessor() *JsonProcessor { _ = "STUB: not implemented"; return nil }

type JsonProcessorVariant interface {
	JsonProcessorCaster() *JsonProcessor
}

func (s *JsonProcessor) JsonProcessorCaster() *JsonProcessor { _ = "STUB: not implemented"; return nil }
