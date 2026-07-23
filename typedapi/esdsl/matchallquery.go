package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _matchAllQuery struct {
	v *types.MatchAllQuery
}

func NewMatchAllQuery() *_matchAllQuery { _ = "STUB: not implemented"; return nil }

func (s *_matchAllQuery) Boost(boost float32) *_matchAllQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchAllQuery) QueryName_(queryname_ string) *_matchAllQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchAllQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_matchAllQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchAllQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchAllQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchAllQuery) MatchAllQueryCaster() *types.MatchAllQuery {
	_ = "STUB: not implemented"
	return nil
}
