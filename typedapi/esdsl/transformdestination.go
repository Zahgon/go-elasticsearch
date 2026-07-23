package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _transformDestination struct {
	v *types.TransformDestination
}

func NewTransformDestination() *_transformDestination { _ = "STUB: not implemented"; return nil }

func (s *_transformDestination) Index(indexname string) *_transformDestination {
	_ = "STUB: not implemented"
	return nil
}

func (s *_transformDestination) Pipeline(pipeline string) *_transformDestination {
	_ = "STUB: not implemented"
	return nil
}

func (s *_transformDestination) TransformDestinationCaster() *types.TransformDestination {
	_ = "STUB: not implemented"
	return nil
}
