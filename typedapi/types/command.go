package types

type Command struct {
	AllocateEmptyPrimary *CommandAllocatePrimaryAction `json:"allocate_empty_primary,omitempty"`

	AllocateReplica *CommandAllocateReplicaAction `json:"allocate_replica,omitempty"`

	AllocateStalePrimary *CommandAllocatePrimaryAction `json:"allocate_stale_primary,omitempty"`

	Cancel *CommandCancelAction `json:"cancel,omitempty"`

	Move *CommandMoveAction `json:"move,omitempty"`
}

func NewCommand() *Command { _ = "STUB: not implemented"; return nil }

type CommandVariant interface {
	CommandCaster() *Command
}

func (s *Command) CommandCaster() *Command { _ = "STUB: not implemented"; return nil }
