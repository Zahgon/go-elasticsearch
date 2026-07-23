package types

type CommandMoveAction struct {
	FromNode string `json:"from_node"`
	Index    string `json:"index"`
	Shard    int    `json:"shard"`

	ToNode string `json:"to_node"`
}

func (s *CommandMoveAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCommandMoveAction() *CommandMoveAction { _ = "STUB: not implemented"; return nil }

type CommandMoveActionVariant interface {
	CommandMoveActionCaster() *CommandMoveAction
}

func (s *CommandMoveAction) CommandMoveActionCaster() *CommandMoveAction {
	_ = "STUB: not implemented"
	return nil
}
