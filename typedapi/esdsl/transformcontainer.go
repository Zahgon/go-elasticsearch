package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _transformContainer struct {
	v *types.TransformContainer
}

func NewTransformContainer() *_transformContainer { _ = "STUB: not implemented"; return nil }

func (s *_transformContainer) Chain(chains ...types.TransformContainerVariant) *_transformContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_transformContainer) ChainValues(chainvalues []types.TransformContainer) *_transformContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_transformContainer) Script(script types.ScriptTransformVariant) *_transformContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_transformContainer) Search(search types.SearchTransformVariant) *_transformContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_transformContainer) TransformContainerCaster() *types.TransformContainer {
	_ = "STUB: not implemented"
	return nil
}
