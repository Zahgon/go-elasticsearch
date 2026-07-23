package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _boolQuery struct {
	v *types.BoolQuery
}

func NewBoolQuery() *_boolQuery { _ = "STUB: not implemented"; return nil }

func (s *_boolQuery) Filter(filters ...types.QueryVariant) *_boolQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boolQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_boolQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boolQuery) Must(musts ...types.QueryVariant) *_boolQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boolQuery) MustNot(mustnots ...types.QueryVariant) *_boolQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boolQuery) Should(shoulds ...types.QueryVariant) *_boolQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boolQuery) Boost(boost float32) *_boolQuery { _ = "STUB: not implemented"; return nil }

func (s *_boolQuery) QueryName_(queryname_ string) *_boolQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boolQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_boolQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boolQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boolQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boolQuery) BoolQueryCaster() *types.BoolQuery { _ = "STUB: not implemented"; return nil }
