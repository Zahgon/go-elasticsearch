package types

type DistanceFeatureQuery any

type DistanceFeatureQueryVariant interface {
	DistanceFeatureQueryCaster() *DistanceFeatureQuery
}
