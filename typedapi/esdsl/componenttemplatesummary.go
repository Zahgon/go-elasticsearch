package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _componentTemplateSummary struct {
	v *types.ComponentTemplateSummary
}

func NewComponentTemplateSummary() *_componentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_componentTemplateSummary) Aliases(aliases map[string]types.AliasDefinition) *_componentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_componentTemplateSummary) AddAlias(key string, value types.AliasDefinitionVariant) *_componentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_componentTemplateSummary) DataStreamOptions(datastreamoptions types.DataStreamOptionsVariant) *_componentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_componentTemplateSummary) Lifecycle(lifecycle types.DataStreamLifecycleVariant) *_componentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_componentTemplateSummary) Mappings(mappings types.TypeMappingVariant) *_componentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_componentTemplateSummary) Meta_(metadata types.MetadataVariant) *_componentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_componentTemplateSummary) Settings(settings map[string]types.IndexSettings) *_componentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_componentTemplateSummary) AddSetting(key string, value types.IndexSettingsVariant) *_componentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_componentTemplateSummary) Version(versionnumber int64) *_componentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_componentTemplateSummary) ComponentTemplateSummaryCaster() *types.ComponentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}
