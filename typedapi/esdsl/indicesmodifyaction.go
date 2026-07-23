package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indicesModifyAction struct {
	v *types.IndicesModifyAction
}

func NewIndicesModifyAction() *_indicesModifyAction { _ = "STUB: not implemented"; return nil }

func (s *_indicesModifyAction) AddBackingIndex(addbackingindex types.IndexAndDataStreamActionVariant) *_indicesModifyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesModifyAction) RemoveBackingIndex(removebackingindex types.IndexAndDataStreamActionVariant) *_indicesModifyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indicesModifyAction) IndicesModifyActionCaster() *types.IndicesModifyAction {
	_ = "STUB: not implemented"
	return nil
}
