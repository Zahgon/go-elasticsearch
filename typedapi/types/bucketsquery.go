package types

type BucketsQuery any

type BucketsQueryVariant interface {
	BucketsQueryCaster() *BucketsQuery
}
