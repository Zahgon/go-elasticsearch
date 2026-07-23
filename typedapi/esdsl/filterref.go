package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/filtertype"
)

type _filterRef struct {
	v *types.FilterRef
}

func NewFilterRef() *_filterRef { _ = "STUB: not implemented"; return nil }

func (s *_filterRef) FilterId(id string) *_filterRef { _ = "STUB: not implemented"; return nil }

func (s *_filterRef) FilterType(filtertype filtertype.FilterType) *_filterRef {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filterRef) FilterRefCaster() *types.FilterRef { _ = "STUB: not implemented"; return nil }
