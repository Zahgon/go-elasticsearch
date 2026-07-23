package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _rankFeatureProperty struct {
	v *types.RankFeatureProperty
}

func NewRankFeatureProperty() *_rankFeatureProperty { _ = "STUB: not implemented"; return nil }

func (s *_rankFeatureProperty) PositiveScoreImpact(positivescoreimpact bool) *_rankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_rankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) Fields(fields map[string]types.Property) *_rankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) AddField(key string, value types.PropertyVariant) *_rankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) IgnoreAbove(ignoreabove int) *_rankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) Meta(meta map[string]string) *_rankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) AddMeta(key string, value string) *_rankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) Properties(properties map[string]types.Property) *_rankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) AddProperty(key string, value types.PropertyVariant) *_rankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_rankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureProperty) RankFeaturePropertyCaster() *types.RankFeatureProperty {
	_ = "STUB: not implemented"
	return nil
}
