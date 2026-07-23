package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _removeAction struct {
	v *types.RemoveAction
}

func NewRemoveAction() *_removeAction { _ = "STUB: not implemented"; return nil }

func (s *_removeAction) Alias(indexalias string) *_removeAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeAction) Aliases(aliases ...string) *_removeAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeAction) Index(indexname string) *_removeAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeAction) Indices(indices ...string) *_removeAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeAction) MustExist(mustexist bool) *_removeAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeAction) IndicesActionCaster() *types.IndicesAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeAction) RemoveActionCaster() *types.RemoveAction {
	_ = "STUB: not implemented"
	return nil
}
