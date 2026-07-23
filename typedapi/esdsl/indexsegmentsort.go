package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortmissing"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortorder"
)

type _indexSegmentSort struct {
	v *types.IndexSegmentSort
}

func NewIndexSegmentSort() *_indexSegmentSort { _ = "STUB: not implemented"; return nil }

func (s *_indexSegmentSort) Field(fields ...string) *_indexSegmentSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSegmentSort) Missing(missings ...segmentsortmissing.SegmentSortMissing) *_indexSegmentSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSegmentSort) Mode(modes ...segmentsortmode.SegmentSortMode) *_indexSegmentSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSegmentSort) Order(orders ...segmentsortorder.SegmentSortOrder) *_indexSegmentSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSegmentSort) IndexSegmentSortCaster() *types.IndexSegmentSort {
	_ = "STUB: not implemented"
	return nil
}
