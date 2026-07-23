package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _searchAccess struct {
	v *types.SearchAccess
}

func NewSearchAccess() *_searchAccess { _ = "STUB: not implemented"; return nil }

func (s *_searchAccess) AllowRestrictedIndices(allowrestrictedindices bool) *_searchAccess {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAccess) FieldSecurity(fieldsecurity types.FieldSecurityVariant) *_searchAccess {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAccess) Names(names ...string) *_searchAccess {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAccess) Query(indicesprivilegesquery types.IndicesPrivilegesQueryVariant) *_searchAccess {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAccess) SearchAccessCaster() *types.SearchAccess {
	_ = "STUB: not implemented"
	return nil
}
