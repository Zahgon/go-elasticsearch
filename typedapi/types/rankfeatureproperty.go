package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type RankFeatureProperty struct {
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	PositiveScoreImpact *bool                                            `json:"positive_score_impact,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *RankFeatureProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s RankFeatureProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRankFeatureProperty() *RankFeatureProperty { _ = "STUB: not implemented"; return nil }

type RankFeaturePropertyVariant interface {
	RankFeaturePropertyCaster() *RankFeatureProperty
}

func (s *RankFeatureProperty) RankFeaturePropertyCaster() *RankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *RankFeatureProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
