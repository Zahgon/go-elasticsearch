package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fielddataFrequencyFilter struct {
	v *types.FielddataFrequencyFilter
}

func NewFielddataFrequencyFilter(max types.Float64, min types.Float64, minsegmentsize int) *_fielddataFrequencyFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fielddataFrequencyFilter) Max(max types.Float64) *_fielddataFrequencyFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fielddataFrequencyFilter) Min(min types.Float64) *_fielddataFrequencyFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fielddataFrequencyFilter) MinSegmentSize(minsegmentsize int) *_fielddataFrequencyFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fielddataFrequencyFilter) FielddataFrequencyFilterCaster() *types.FielddataFrequencyFilter {
	_ = "STUB: not implemented"
	return nil
}
