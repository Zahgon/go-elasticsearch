package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rangeQuery struct {
	v types.RangeQuery
}

func NewRangeQuery() *_rangeQuery { _ = "STUB: not implemented"; return nil }

func (u *_rangeQuery) UntypedRangeQuery(untypedrangequery types.UntypedRangeQueryVariant) *_rangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_untypedRangeQuery) RangeQueryCaster() *types.RangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rangeQuery) DateRangeQuery(daterangequery types.DateRangeQueryVariant) *_rangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_dateRangeQuery) RangeQueryCaster() *types.RangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rangeQuery) NumberRangeQuery(numberrangequery types.NumberRangeQueryVariant) *_rangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_numberRangeQuery) RangeQueryCaster() *types.RangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rangeQuery) LongNumberRangeQuery(longnumberrangequery types.LongNumberRangeQueryVariant) *_rangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_longNumberRangeQuery) RangeQueryCaster() *types.RangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rangeQuery) TermRangeQuery(termrangequery types.TermRangeQueryVariant) *_rangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_termRangeQuery) RangeQueryCaster() *types.RangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rangeQuery) RangeQueryCaster() *types.RangeQuery { _ = "STUB: not implemented"; return nil }
