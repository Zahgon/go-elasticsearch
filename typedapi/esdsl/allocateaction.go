package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _allocateAction struct {
	v *types.AllocateAction
}

func NewAllocateAction() *_allocateAction { _ = "STUB: not implemented"; return nil }

func (s *_allocateAction) Exclude(exclude map[string]string) *_allocateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allocateAction) AddExclude(key string, value string) *_allocateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allocateAction) Include(include map[string]string) *_allocateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allocateAction) AddInclude(key string, value string) *_allocateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allocateAction) NumberOfReplicas(numberofreplicas int) *_allocateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allocateAction) Require(require map[string]string) *_allocateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allocateAction) AddRequire(key string, value string) *_allocateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allocateAction) TotalShardsPerNode(totalshardspernode int) *_allocateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allocateAction) AllocateActionCaster() *types.AllocateAction {
	_ = "STUB: not implemented"
	return nil
}
