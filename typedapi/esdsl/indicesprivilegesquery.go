package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indicesPrivilegesQuery struct {
	v types.IndicesPrivilegesQuery
}

func NewIndicesPrivilegesQuery() *_indicesPrivilegesQuery { _ = "STUB: not implemented"; return nil }

func (u *_indicesPrivilegesQuery) String(string string) *_indicesPrivilegesQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_indicesPrivilegesQuery) Query(query types.QueryVariant) *_indicesPrivilegesQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_query) IndicesPrivilegesQueryCaster() *types.IndicesPrivilegesQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_indicesPrivilegesQuery) RoleTemplateQuery(roletemplatequery types.RoleTemplateQueryVariant) *_indicesPrivilegesQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_roleTemplateQuery) IndicesPrivilegesQueryCaster() *types.IndicesPrivilegesQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_indicesPrivilegesQuery) IndicesPrivilegesQueryCaster() *types.IndicesPrivilegesQuery {
	_ = "STUB: not implemented"
	return nil
}
