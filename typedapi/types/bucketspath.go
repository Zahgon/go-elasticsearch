package types

type BucketsPath any

type BucketsPathVariant interface {
	BucketsPathCaster() *BucketsPath
}
