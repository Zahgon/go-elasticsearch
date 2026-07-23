package types

type Queries struct {
	Cache *CacheQueries `json:"cache,omitempty"`
}

func NewQueries() *Queries { _ = "STUB: not implemented"; return nil }

type QueriesVariant interface {
	QueriesCaster() *Queries
}

func (s *Queries) QueriesCaster() *Queries { _ = "STUB: not implemented"; return nil }
