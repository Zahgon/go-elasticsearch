package types

type SlackAction struct {
	Account *string      `json:"account,omitempty"`
	Message SlackMessage `json:"message"`
}

func (s *SlackAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSlackAction() *SlackAction { _ = "STUB: not implemented"; return nil }

type SlackActionVariant interface {
	SlackActionCaster() *SlackAction
}

func (s *SlackAction) SlackActionCaster() *SlackAction { _ = "STUB: not implemented"; return nil }
