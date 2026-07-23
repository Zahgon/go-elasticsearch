package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/subobjects"
)

type _typeMapping struct {
	v *types.TypeMapping
}

func NewTypeMapping() *_typeMapping { _ = "STUB: not implemented"; return nil }

func (s *_typeMapping) AllField(allfield types.AllFieldVariant) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) DataStreamTimestamp_(datastreamtimestamp_ types.DataStreamTimestampVariant) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) DateDetection(datedetection bool) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) Dynamic(dynamic dynamicmapping.DynamicMapping) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) DynamicDateFormats(dynamicdateformats ...string) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) DynamicTemplates(dynamictemplates []map[string]types.DynamicTemplate) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) Enabled(enabled bool) *_typeMapping { _ = "STUB: not implemented"; return nil }

func (s *_typeMapping) FieldNames_(fieldnames_ types.FieldNamesFieldVariant) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) IndexField(indexfield types.IndexFieldVariant) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) Meta_(metadata types.MetadataVariant) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) NumericDetection(numericdetection bool) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) Properties(properties map[string]types.Property) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) AddProperty(key string, value types.PropertyVariant) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) Routing_(routing_ types.RoutingFieldVariant) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) Runtime(runtime map[string]types.RuntimeField) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) AddRuntime(key string, value types.RuntimeFieldVariant) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) Size_(size_ types.SizeFieldVariant) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) Source_(source_ types.SourceFieldVariant) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) Subobjects(subobjects subobjects.Subobjects) *_typeMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_typeMapping) TypeMappingCaster() *types.TypeMapping {
	_ = "STUB: not implemented"
	return nil
}
