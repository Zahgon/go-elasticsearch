package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _idsQuery struct {
	v *types.IdsQuery
}

func NewIdsQuery() *_idsQuery { _ = "STUB: not implemented"; return nil }

func (s *_idsQuery) Values(ids ...string) *_idsQuery { _ = "STUB: not implemented"; return nil }

func (s *_idsQuery) Boost(boost float32) *_idsQuery { _ = "STUB: not implemented"; return nil }

func (s *_idsQuery) QueryName_(queryname_ string) *_idsQuery { _ = "STUB: not implemented"; return nil }

func (s *_idsQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_idsQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_idsQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_idsQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_idsQuery) IdsQueryCaster() *types.IdsQuery { _ = "STUB: not implemented"; return nil }
