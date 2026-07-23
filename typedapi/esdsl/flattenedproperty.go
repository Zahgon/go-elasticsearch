package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _flattenedProperty struct {
	v *types.FlattenedProperty
}

func NewFlattenedProperty() *_flattenedProperty { _ = "STUB: not implemented"; return nil }

func (s *_flattenedProperty) Boost(boost types.Float64) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) DepthLimit(depthlimit int) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) DocValues(docvalues bool) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) EagerGlobalOrdinals(eagerglobalordinals bool) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) Index(index bool) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) NullValue(nullvalue string) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) Similarity(similarity string) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) SplitQueriesOnWhitespace(splitqueriesonwhitespace bool) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) TimeSeriesDimensions(timeseriesdimensions ...string) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) Fields(fields map[string]types.Property) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) AddField(key string, value types.PropertyVariant) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) IgnoreAbove(ignoreabove int) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) Meta(meta map[string]string) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) AddMeta(key string, value string) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) Properties(properties map[string]types.Property) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) AddProperty(key string, value types.PropertyVariant) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_flattenedProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_flattenedProperty) FlattenedPropertyCaster() *types.FlattenedProperty {
	_ = "STUB: not implemented"
	return nil
}
