package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pinnedQuery struct {
	v *types.PinnedQuery
}

func NewPinnedQuery() *_pinnedQuery { _ = "STUB: not implemented"; return nil }

func (s *_pinnedQuery) Docs(docs ...types.PinnedDocVariant) *_pinnedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedQuery) DocsValues(docsvalues []types.PinnedDoc) *_pinnedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedQuery) Ids(ids ...string) *_pinnedQuery { _ = "STUB: not implemented"; return nil }

func (s *_pinnedQuery) Organic(organic types.QueryVariant) *_pinnedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedQuery) Boost(boost float32) *_pinnedQuery { _ = "STUB: not implemented"; return nil }

func (s *_pinnedQuery) QueryName_(queryname_ string) *_pinnedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedQuery) PinnedQueryCaster() *types.PinnedQuery {
	_ = "STUB: not implemented"
	return nil
}
