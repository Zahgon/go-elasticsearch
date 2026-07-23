package types

type CommandCancelAction struct {
	AllowPrimary *bool  `json:"allow_primary,omitempty"`
	Index        string `json:"index"`
	Node         string `json:"node"`
	Shard        int    `json:"shard"`
}

func (s *CommandCancelAction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCommandCancelAction() *CommandCancelAction { _ = "STUB: not implemented"; return nil }

type CommandCancelActionVariant interface {
	CommandCancelActionCaster() *CommandCancelAction
}

func (s *CommandCancelAction) CommandCancelActionCaster() *CommandCancelAction {
	_ = "STUB: not implemented"
	return nil
}
