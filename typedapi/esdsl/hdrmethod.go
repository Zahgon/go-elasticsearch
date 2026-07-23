package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hdrMethod struct {
	v *types.HdrMethod
}

func NewHdrMethod() *_hdrMethod { _ = "STUB: not implemented"; return nil }

func (s *_hdrMethod) NumberOfSignificantValueDigits(numberofsignificantvaluedigits int) *_hdrMethod {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hdrMethod) HdrMethodCaster() *types.HdrMethod { _ = "STUB: not implemented"; return nil }
