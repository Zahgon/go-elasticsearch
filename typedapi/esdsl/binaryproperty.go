package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _binaryProperty struct {
	v *types.BinaryProperty
}

func NewBinaryProperty() *_binaryProperty { _ = "STUB: not implemented"; return nil }

func (s *_binaryProperty) CopyTo(fields ...string) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) DocValues(docvalues bool) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) Fields(fields map[string]types.Property) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) AddField(key string, value types.PropertyVariant) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) IgnoreAbove(ignoreabove int) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) Meta(meta map[string]string) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) AddMeta(key string, value string) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) Properties(properties map[string]types.Property) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) AddProperty(key string, value types.PropertyVariant) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) Store(store bool) *_binaryProperty { _ = "STUB: not implemented"; return nil }

func (s *_binaryProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_binaryProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_binaryProperty) BinaryPropertyCaster() *types.BinaryProperty {
	_ = "STUB: not implemented"
	return nil
}
