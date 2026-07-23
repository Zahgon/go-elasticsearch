package types

type CommandAllocateReplicaAction struct {
	Index string `json:"index"`
	Node  string `json:"node"`
	Shard int    `json:"shard"`
}

func (s *CommandAllocateReplicaAction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCommandAllocateReplicaAction() *CommandAllocateReplicaAction {
	_ = "STUB: not implemented"
	return nil
}

type CommandAllocateReplicaActionVariant interface {
	CommandAllocateReplicaActionCaster() *CommandAllocateReplicaAction
}

func (s *CommandAllocateReplicaAction) CommandAllocateReplicaActionCaster() *CommandAllocateReplicaAction {
	_ = "STUB: not implemented"
	return nil
}
