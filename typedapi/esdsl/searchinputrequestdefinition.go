package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

type _searchInputRequestDefinition struct {
	v *types.SearchInputRequestDefinition
}

func NewSearchInputRequestDefinition() *_searchInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInputRequestDefinition) Body(body types.SearchInputRequestBodyVariant) *_searchInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInputRequestDefinition) Indices(indices ...string) *_searchInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInputRequestDefinition) IndicesOptions(indicesoptions types.IndicesOptionsVariant) *_searchInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInputRequestDefinition) RestTotalHitsAsInt(resttotalhitsasint bool) *_searchInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInputRequestDefinition) SearchType(searchtype searchtype.SearchType) *_searchInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInputRequestDefinition) Template(template types.SearchTemplateRequestBodyVariant) *_searchInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInputRequestDefinition) SearchInputRequestDefinitionCaster() *types.SearchInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}
