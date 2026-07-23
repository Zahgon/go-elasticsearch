package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _minHashTokenFilter struct {
	v *types.MinHashTokenFilter
}

func NewMinHashTokenFilter() *_minHashTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_minHashTokenFilter) BucketCount(bucketcount int) *_minHashTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minHashTokenFilter) HashCount(hashcount int) *_minHashTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minHashTokenFilter) HashSetSize(hashsetsize int) *_minHashTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minHashTokenFilter) WithRotation(withrotation bool) *_minHashTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minHashTokenFilter) Version(versionstring string) *_minHashTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minHashTokenFilter) MinHashTokenFilterCaster() *types.MinHashTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
