package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexTemplateMapping struct {
	v *types.IndexTemplateMapping
}

func NewIndexTemplateMapping() *_indexTemplateMapping { _ = "STUB: not implemented"; return nil }

func (s *_indexTemplateMapping) Aliases(aliases map[string]types.Alias) *_indexTemplateMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateMapping) AddAlias(key string, value types.AliasVariant) *_indexTemplateMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateMapping) DataStreamOptions(datastreamoptions types.DataStreamOptionsTemplateVariant) *_indexTemplateMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateMapping) Lifecycle(lifecycle types.DataStreamLifecycleVariant) *_indexTemplateMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateMapping) Mappings(mappings types.TypeMappingVariant) *_indexTemplateMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateMapping) Settings(settings types.IndexSettingsVariant) *_indexTemplateMapping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateMapping) IndexTemplateMappingCaster() *types.IndexTemplateMapping {
	_ = "STUB: not implemented"
	return nil
}
