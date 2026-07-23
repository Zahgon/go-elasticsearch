package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _passthroughObjectProperty struct {
	v *types.PassthroughObjectProperty
}

func NewPassthroughObjectProperty() *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) Enabled(enabled bool) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) Priority(priority int) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) TimeSeriesDimension(timeseriesdimension bool) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) CopyTo(fields ...string) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) Fields(fields map[string]types.Property) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) AddField(key string, value types.PropertyVariant) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) IgnoreAbove(ignoreabove int) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) Meta(meta map[string]string) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) AddMeta(key string, value string) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) Properties(properties map[string]types.Property) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) AddProperty(key string, value types.PropertyVariant) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) Store(store bool) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_passthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_passthroughObjectProperty) PassthroughObjectPropertyCaster() *types.PassthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}
