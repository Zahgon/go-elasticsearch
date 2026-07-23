package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _createFrom struct {
	v *types.CreateFrom
}

func NewCreateFrom() *_createFrom { _ = "STUB: not implemented"; return nil }

func (s *_createFrom) MappingsOverride(mappingsoverride types.TypeMappingVariant) *_createFrom {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createFrom) RemoveIndexBlocks(removeindexblocks bool) *_createFrom {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createFrom) SettingsOverride(settingsoverride types.IndexSettingsVariant) *_createFrom {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createFrom) CreateFromCaster() *types.CreateFrom { _ = "STUB: not implemented"; return nil }
