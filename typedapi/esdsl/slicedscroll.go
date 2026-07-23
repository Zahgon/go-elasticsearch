package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slicedScroll struct {
	v *types.SlicedScroll
}

func NewSlicedScroll(max int) *_slicedScroll { _ = "STUB: not implemented"; return nil }

func (s *_slicedScroll) Field(field string) *_slicedScroll { _ = "STUB: not implemented"; return nil }

func (s *_slicedScroll) Id(id string) *_slicedScroll { _ = "STUB: not implemented"; return nil }

func (s *_slicedScroll) Max(max int) *_slicedScroll { _ = "STUB: not implemented"; return nil }

func (s *_slicedScroll) SlicedScrollCaster() *types.SlicedScroll {
	_ = "STUB: not implemented"
	return nil
}
