package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _mappingCharFilter struct {
	v *types.MappingCharFilter
}

func NewMappingCharFilter() *_mappingCharFilter { _ = "STUB: not implemented"; return nil }

func (s *_mappingCharFilter) Mappings(mappings ...string) *_mappingCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingCharFilter) MappingsPath(mappingspath string) *_mappingCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingCharFilter) Version(versionstring string) *_mappingCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingCharFilter) MappingCharFilterCaster() *types.MappingCharFilter {
	_ = "STUB: not implemented"
	return nil
}
