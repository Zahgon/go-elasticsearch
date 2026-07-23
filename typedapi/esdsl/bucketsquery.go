package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bucketsQuery struct {
	v types.BucketsQuery
}

func NewBucketsQuery() *_bucketsQuery { _ = "STUB: not implemented"; return nil }

func (u *_bucketsQuery) Map(value map[string]types.QueryVariant) *_bucketsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_bucketsQuery) QueryContainers(querycontainers ...types.QueryVariant) *_bucketsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_bucketsQuery) BucketsQueryCaster() *types.BucketsQuery {
	_ = "STUB: not implemented"
	return nil
}
