package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type RankFeaturesProperty struct {
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	PositiveScoreImpact *bool                                            `json:"positive_score_impact,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *RankFeaturesProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s RankFeaturesProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRankFeaturesProperty() *RankFeaturesProperty { _ = "STUB: not implemented"; return nil }

type RankFeaturesPropertyVariant interface {
	RankFeaturesPropertyCaster() *RankFeaturesProperty
}

func (s *RankFeaturesProperty) RankFeaturesPropertyCaster() *RankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *RankFeaturesProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
