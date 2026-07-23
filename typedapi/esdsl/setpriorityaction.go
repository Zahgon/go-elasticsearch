package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _setPriorityAction struct {
	v *types.SetPriorityAction
}

func NewSetPriorityAction() *_setPriorityAction { _ = "STUB: not implemented"; return nil }

func (s *_setPriorityAction) Priority(priority int) *_setPriorityAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setPriorityAction) SetPriorityActionCaster() *types.SetPriorityAction {
	_ = "STUB: not implemented"
	return nil
}
