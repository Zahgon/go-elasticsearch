package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _completionContext struct {
	v *types.CompletionContext
}

func NewCompletionContext() *_completionContext { _ = "STUB: not implemented"; return nil }

func (s *_completionContext) Boost(boost types.Float64) *_completionContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionContext) Context(context types.ContextVariant) *_completionContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionContext) Neighbours(neighbours ...types.GeoHashPrecisionVariant) *_completionContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionContext) NeighboursValues(neighboursvalues []types.GeoHashPrecision) *_completionContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionContext) Precision(geohashprecision types.GeoHashPrecisionVariant) *_completionContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionContext) Prefix(prefix bool) *_completionContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionContext) CompletionContextCaster() *types.CompletionContext {
	_ = "STUB: not implemented"
	return nil
}
