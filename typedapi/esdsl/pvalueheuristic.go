package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pValueHeuristic struct {
	v *types.PValueHeuristic
}

func NewPValueHeuristic() *_pValueHeuristic { _ = "STUB: not implemented"; return nil }

func (s *_pValueHeuristic) BackgroundIsSuperset(backgroundissuperset bool) *_pValueHeuristic {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pValueHeuristic) NormalizeAbove(normalizeabove int64) *_pValueHeuristic {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pValueHeuristic) PValueHeuristicCaster() *types.PValueHeuristic {
	_ = "STUB: not implemented"
	return nil
}
