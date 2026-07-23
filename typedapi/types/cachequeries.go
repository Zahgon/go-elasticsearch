package types

type CacheQueries struct {
	Enabled bool `json:"enabled"`
}

func (s *CacheQueries) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCacheQueries() *CacheQueries { _ = "STUB: not implemented"; return nil }

type CacheQueriesVariant interface {
	CacheQueriesCaster() *CacheQueries
}

func (s *CacheQueries) CacheQueriesCaster() *CacheQueries { _ = "STUB: not implemented"; return nil }
