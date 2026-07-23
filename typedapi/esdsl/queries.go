package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _queries struct {
	v *types.Queries
}

func NewQueries() *_queries { _ = "STUB: not implemented"; return nil }

func (s *_queries) Cache(cache types.CacheQueriesVariant) *_queries {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queries) QueriesCaster() *types.Queries { _ = "STUB: not implemented"; return nil }
