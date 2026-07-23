package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dateDistanceFeatureQuery struct {
	v *types.DateDistanceFeatureQuery
}

func NewDateDistanceFeatureQuery() *_dateDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDistanceFeatureQuery) Boost(boost float32) *_dateDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDistanceFeatureQuery) Field(field string) *_dateDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDistanceFeatureQuery) Origin(datemath string) *_dateDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDistanceFeatureQuery) Pivot(duration types.DurationVariant) *_dateDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDistanceFeatureQuery) QueryName_(queryname_ string) *_dateDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDistanceFeatureQuery) QueryCaster() *types.Query {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateDistanceFeatureQuery) DateDistanceFeatureQueryCaster() *types.DateDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}
