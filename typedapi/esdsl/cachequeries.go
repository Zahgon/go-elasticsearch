package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _cacheQueries struct {
	v *types.CacheQueries
}

func NewCacheQueries(enabled bool) *_cacheQueries { _ = "STUB: not implemented"; return nil }

func (s *_cacheQueries) Enabled(enabled bool) *_cacheQueries { _ = "STUB: not implemented"; return nil }

func (s *_cacheQueries) CacheQueriesCaster() *types.CacheQueries {
	_ = "STUB: not implemented"
	return nil
}
