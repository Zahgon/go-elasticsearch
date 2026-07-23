package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _command struct {
	v *types.Command
}

func NewCommand() *_command { _ = "STUB: not implemented"; return nil }

func (s *_command) AllocateEmptyPrimary(allocateemptyprimary types.CommandAllocatePrimaryActionVariant) *_command {
	_ = "STUB: not implemented"
	return nil
}

func (s *_command) AllocateReplica(allocatereplica types.CommandAllocateReplicaActionVariant) *_command {
	_ = "STUB: not implemented"
	return nil
}

func (s *_command) AllocateStalePrimary(allocatestaleprimary types.CommandAllocatePrimaryActionVariant) *_command {
	_ = "STUB: not implemented"
	return nil
}

func (s *_command) Cancel(cancel types.CommandCancelActionVariant) *_command {
	_ = "STUB: not implemented"
	return nil
}

func (s *_command) Move(move types.CommandMoveActionVariant) *_command {
	_ = "STUB: not implemented"
	return nil
}

func (s *_command) CommandCaster() *types.Command { _ = "STUB: not implemented"; return nil }
