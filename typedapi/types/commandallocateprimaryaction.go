package types

type CommandAllocatePrimaryAction struct {
	AcceptDataLoss bool   `json:"accept_data_loss"`
	Index          string `json:"index"`
	Node           string `json:"node"`
	Shard          int    `json:"shard"`
}

func (s *CommandAllocatePrimaryAction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCommandAllocatePrimaryAction() *CommandAllocatePrimaryAction {
	_ = "STUB: not implemented"
	return nil
}

type CommandAllocatePrimaryActionVariant interface {
	CommandAllocatePrimaryActionCaster() *CommandAllocatePrimaryAction
}

func (s *CommandAllocatePrimaryAction) CommandAllocatePrimaryActionCaster() *CommandAllocatePrimaryAction {
	_ = "STUB: not implemented"
	return nil
}
