package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _commandAllocateReplicaAction struct {
	v *types.CommandAllocateReplicaAction
}

func NewCommandAllocateReplicaAction(node string, shard int) *_commandAllocateReplicaAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandAllocateReplicaAction) Index(indexname string) *_commandAllocateReplicaAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandAllocateReplicaAction) Node(node string) *_commandAllocateReplicaAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandAllocateReplicaAction) Shard(shard int) *_commandAllocateReplicaAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandAllocateReplicaAction) CommandAllocateReplicaActionCaster() *types.CommandAllocateReplicaAction {
	_ = "STUB: not implemented"
	return nil
}
