package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoshaperelation"
)

type EnrichProcessor struct {
	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	MaxMatches *int `json:"max_matches,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Override *bool `json:"override,omitempty"`

	PolicyName string `json:"policy_name"`

	ShapeRelation *geoshaperelation.GeoShapeRelation `json:"shape_relation,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField string `json:"target_field"`
}

func (s *EnrichProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEnrichProcessor() *EnrichProcessor { _ = "STUB: not implemented"; return nil }

type EnrichProcessorVariant interface {
	EnrichProcessorCaster() *EnrichProcessor
}

func (s *EnrichProcessor) EnrichProcessorCaster() *EnrichProcessor {
	_ = "STUB: not implemented"
	return nil
}
