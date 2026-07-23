package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/filtertype"
)

type FilterRef struct {
	FilterId string `json:"filter_id"`

	FilterType *filtertype.FilterType `json:"filter_type,omitempty"`
}

func (s *FilterRef) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFilterRef() *FilterRef { _ = "STUB: not implemented"; return nil }

type FilterRefVariant interface {
	FilterRefCaster() *FilterRef
}

func (s *FilterRef) FilterRefCaster() *FilterRef { _ = "STUB: not implemented"; return nil }
