package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indicesAction struct {
	v *types.IndicesAction
}

func NewIndicesAction() *_indicesAction { _ = "STUB: not implemented"; return nil }

func (s *_indicesAction) Add(add types.AddActionVariant) *_indicesAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesAction) Remove(remove types.RemoveActionVariant) *_indicesAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesAction) RemoveIndex(removeindex types.RemoveIndexActionVariant) *_indicesAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesAction) IndicesActionCaster() *types.IndicesAction {
	_ = "STUB: not implemented"
	return nil
}
