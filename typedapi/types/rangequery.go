package types

type RangeQuery any

type RangeQueryVariant interface {
	RangeQueryCaster() *RangeQuery
}
