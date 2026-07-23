package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _percolateQuery struct {
	v *types.PercolateQuery
}

func NewPercolateQuery() *_percolateQuery { _ = "STUB: not implemented"; return nil }

func (s *_percolateQuery) Document(document json.RawMessage) *_percolateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolateQuery) Documents(documents ...json.RawMessage) *_percolateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolateQuery) Field(field string) *_percolateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolateQuery) Id(id string) *_percolateQuery { _ = "STUB: not implemented"; return nil }

func (s *_percolateQuery) Index(indexname string) *_percolateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolateQuery) Name(name string) *_percolateQuery { _ = "STUB: not implemented"; return nil }

func (s *_percolateQuery) Preference(preference string) *_percolateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolateQuery) Routing(routing string) *_percolateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolateQuery) Version(versionnumber int64) *_percolateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolateQuery) Boost(boost float32) *_percolateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolateQuery) QueryName_(queryname_ string) *_percolateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percolateQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_percolateQuery) PercolateQueryCaster() *types.PercolateQuery {
	_ = "STUB: not implemented"
	return nil
}
