package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _searchTransform struct {
	v *types.SearchTransform
}

func NewSearchTransform(request types.SearchInputRequestDefinitionVariant) *_searchTransform {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTransform) Request(request types.SearchInputRequestDefinitionVariant) *_searchTransform {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTransform) Timeout(duration types.DurationVariant) *_searchTransform {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTransform) TransformContainerCaster() *types.TransformContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTransform) SearchTransformCaster() *types.SearchTransform {
	_ = "STUB: not implemented"
	return nil
}
