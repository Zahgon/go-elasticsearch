package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _removeIndexAction struct {
	v *types.RemoveIndexAction
}

func NewRemoveIndexAction() *_removeIndexAction { _ = "STUB: not implemented"; return nil }

func (s *_removeIndexAction) Index(indexname string) *_removeIndexAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeIndexAction) Indices(indices ...string) *_removeIndexAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeIndexAction) MustExist(mustexist bool) *_removeIndexAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeIndexAction) IndicesActionCaster() *types.IndicesAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeIndexAction) RemoveIndexActionCaster() *types.RemoveIndexAction {
	_ = "STUB: not implemented"
	return nil
}
