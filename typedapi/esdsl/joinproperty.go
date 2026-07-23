package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _joinProperty struct {
	v *types.JoinProperty
}

func NewJoinProperty() *_joinProperty { _ = "STUB: not implemented"; return nil }

func (s *_joinProperty) EagerGlobalOrdinals(eagerglobalordinals bool) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) Relations(relations map[string][]string) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) Fields(fields map[string]types.Property) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) AddField(key string, value types.PropertyVariant) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) IgnoreAbove(ignoreabove int) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) Meta(meta map[string]string) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) AddMeta(key string, value string) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) Properties(properties map[string]types.Property) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) AddProperty(key string, value types.PropertyVariant) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_joinProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProperty) JoinPropertyCaster() *types.JoinProperty {
	_ = "STUB: not implemented"
	return nil
}
