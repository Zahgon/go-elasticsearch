package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rescoreVector struct {
	v *types.RescoreVector
}

func NewRescoreVector(oversample float32) *_rescoreVector { _ = "STUB: not implemented"; return nil }

func (s *_rescoreVector) Oversample(oversample float32) *_rescoreVector {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescoreVector) RescoreVectorCaster() *types.RescoreVector {
	_ = "STUB: not implemented"
	return nil
}
