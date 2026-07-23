package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortmissing"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/segmentsortorder"
)

type IndexSegmentSort struct {
	Field   []string                                `json:"field,omitempty"`
	Missing []segmentsortmissing.SegmentSortMissing `json:"missing,omitempty"`
	Mode    []segmentsortmode.SegmentSortMode       `json:"mode,omitempty"`
	Order   []segmentsortorder.SegmentSortOrder     `json:"order,omitempty"`
}

func (s *IndexSegmentSort) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexSegmentSort() *IndexSegmentSort { _ = "STUB: not implemented"; return nil }

type IndexSegmentSortVariant interface {
	IndexSegmentSortCaster() *IndexSegmentSort
}

func (s *IndexSegmentSort) IndexSegmentSortCaster() *IndexSegmentSort {
	_ = "STUB: not implemented"
	return nil
}
