package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexTemplateSummary struct {
	v *types.IndexTemplateSummary
}

func NewIndexTemplateSummary() *_indexTemplateSummary { _ = "STUB: not implemented"; return nil }

func (s *_indexTemplateSummary) Aliases(aliases map[string]types.Alias) *_indexTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateSummary) AddAlias(key string, value types.AliasVariant) *_indexTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateSummary) DataStreamOptions(datastreamoptions types.DataStreamOptionsVariant) *_indexTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateSummary) Lifecycle(lifecycle types.DataStreamLifecycleVariant) *_indexTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateSummary) Mappings(mappings types.TypeMappingVariant) *_indexTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateSummary) Settings(settings types.IndexSettingsVariant) *_indexTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplateSummary) IndexTemplateSummaryCaster() *types.IndexTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}
