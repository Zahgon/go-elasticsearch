package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/excludefrequent"
)

type Detector struct {
	ByFieldName *string `json:"by_field_name,omitempty"`

	CustomRules []DetectionRule `json:"custom_rules,omitempty"`

	DetectorDescription *string `json:"detector_description,omitempty"`

	DetectorIndex *int `json:"detector_index,omitempty"`

	ExcludeFrequent *excludefrequent.ExcludeFrequent `json:"exclude_frequent,omitempty"`

	FieldName *string `json:"field_name,omitempty"`

	Function *string `json:"function,omitempty"`

	OverFieldName *string `json:"over_field_name,omitempty"`

	PartitionFieldName *string `json:"partition_field_name,omitempty"`

	UseNull *bool `json:"use_null,omitempty"`
}

func (s *Detector) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDetector() *Detector { _ = "STUB: not implemented"; return nil }

type DetectorVariant interface {
	DetectorCaster() *Detector
}

func (s *Detector) DetectorCaster() *Detector { _ = "STUB: not implemented"; return nil }
