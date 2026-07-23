package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _matchOnlyTextProperty struct {
	v *types.MatchOnlyTextProperty
}

func NewMatchOnlyTextProperty() *_matchOnlyTextProperty { _ = "STUB: not implemented"; return nil }

func (s *_matchOnlyTextProperty) CopyTo(fields ...string) *_matchOnlyTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchOnlyTextProperty) Fields(fields map[string]types.Property) *_matchOnlyTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchOnlyTextProperty) AddField(key string, value types.PropertyVariant) *_matchOnlyTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchOnlyTextProperty) Meta(meta map[string]string) *_matchOnlyTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchOnlyTextProperty) AddMeta(key string, value string) *_matchOnlyTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchOnlyTextProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchOnlyTextProperty) MatchOnlyTextPropertyCaster() *types.MatchOnlyTextProperty {
	_ = "STUB: not implemented"
	return nil
}
