package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _searchInputRequestBody struct {
	v *types.SearchInputRequestBody
}

func NewSearchInputRequestBody(query types.QueryVariant) *_searchInputRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInputRequestBody) Query(query types.QueryVariant) *_searchInputRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchInputRequestBody) SearchInputRequestBodyCaster() *types.SearchInputRequestBody {
	_ = "STUB: not implemented"
	return nil
}
