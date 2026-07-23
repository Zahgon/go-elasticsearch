package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _rankFeaturesProperty struct {
	v *types.RankFeaturesProperty
}

func NewRankFeaturesProperty() *_rankFeaturesProperty { _ = "STUB: not implemented"; return nil }

func (s *_rankFeaturesProperty) PositiveScoreImpact(positivescoreimpact bool) *_rankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_rankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) Fields(fields map[string]types.Property) *_rankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) AddField(key string, value types.PropertyVariant) *_rankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) IgnoreAbove(ignoreabove int) *_rankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) Meta(meta map[string]string) *_rankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) AddMeta(key string, value string) *_rankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) Properties(properties map[string]types.Property) *_rankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) AddProperty(key string, value types.PropertyVariant) *_rankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_rankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeaturesProperty) RankFeaturesPropertyCaster() *types.RankFeaturesProperty {
	_ = "STUB: not implemented"
	return nil
}
