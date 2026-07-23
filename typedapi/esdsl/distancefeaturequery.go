package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _distanceFeatureQuery struct {
	v types.DistanceFeatureQuery
}

func NewDistanceFeatureQuery() *_distanceFeatureQuery { _ = "STUB: not implemented"; return nil }

func (u *_distanceFeatureQuery) UntypedDistanceFeatureQuery(untypeddistancefeaturequery types.UntypedDistanceFeatureQueryVariant) *_distanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_untypedDistanceFeatureQuery) DistanceFeatureQueryCaster() *types.DistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_distanceFeatureQuery) GeoDistanceFeatureQuery(geodistancefeaturequery types.GeoDistanceFeatureQueryVariant) *_distanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoDistanceFeatureQuery) DistanceFeatureQueryCaster() *types.DistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_distanceFeatureQuery) DateDistanceFeatureQuery(datedistancefeaturequery types.DateDistanceFeatureQueryVariant) *_distanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_dateDistanceFeatureQuery) DistanceFeatureQueryCaster() *types.DistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (u *_distanceFeatureQuery) DistanceFeatureQueryCaster() *types.DistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}
