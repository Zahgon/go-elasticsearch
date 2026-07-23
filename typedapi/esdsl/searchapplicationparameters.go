package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _searchApplicationParameters struct {
	v *types.SearchApplicationParameters
}

func NewSearchApplicationParameters() *_searchApplicationParameters {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchApplicationParameters) AnalyticsCollectionName(name string) *_searchApplicationParameters {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchApplicationParameters) Indices(indices ...string) *_searchApplicationParameters {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchApplicationParameters) Template(template types.SearchApplicationTemplateVariant) *_searchApplicationParameters {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchApplicationParameters) SearchApplicationParametersCaster() *types.SearchApplicationParameters {
	_ = "STUB: not implemented"
	return nil
}
