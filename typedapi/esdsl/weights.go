package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _weights struct {
	v *types.Weights
}

func NewWeights(weights types.Float64) *_weights { _ = "STUB: not implemented"; return nil }

func (s *_weights) Weights(weights types.Float64) *_weights { _ = "STUB: not implemented"; return nil }

func (s *_weights) WeightsCaster() *types.Weights { _ = "STUB: not implemented"; return nil }
