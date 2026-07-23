package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _commandAllocatePrimaryAction struct {
	v *types.CommandAllocatePrimaryAction
}

func NewCommandAllocatePrimaryAction(acceptdataloss bool, node string, shard int) *_commandAllocatePrimaryAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandAllocatePrimaryAction) AcceptDataLoss(acceptdataloss bool) *_commandAllocatePrimaryAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandAllocatePrimaryAction) Index(indexname string) *_commandAllocatePrimaryAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandAllocatePrimaryAction) Node(node string) *_commandAllocatePrimaryAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandAllocatePrimaryAction) Shard(shard int) *_commandAllocatePrimaryAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commandAllocatePrimaryAction) CommandAllocatePrimaryActionCaster() *types.CommandAllocatePrimaryAction {
	_ = "STUB: not implemented"
	return nil
}
