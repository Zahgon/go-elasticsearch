package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _semanticTextProperty struct {
	v *types.SemanticTextProperty
}

func NewSemanticTextProperty() *_semanticTextProperty { _ = "STUB: not implemented"; return nil }

func (s *_semanticTextProperty) ChunkingSettings(chunkingsettings types.ChunkingSettingsVariant) *_semanticTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextProperty) Fields(fields map[string]types.Property) *_semanticTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextProperty) AddField(key string, value types.PropertyVariant) *_semanticTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextProperty) IndexOptions(indexoptions types.SemanticTextIndexOptionsVariant) *_semanticTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextProperty) InferenceId(id string) *_semanticTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextProperty) Meta(meta map[string]string) *_semanticTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextProperty) AddMeta(key string, value string) *_semanticTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextProperty) SearchInferenceId(id string) *_semanticTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextProperty) SemanticTextPropertyCaster() *types.SemanticTextProperty {
	_ = "STUB: not implemented"
	return nil
}
