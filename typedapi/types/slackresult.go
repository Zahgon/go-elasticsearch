package types

type SlackResult struct {
	Account *string      `json:"account,omitempty"`
	Message SlackMessage `json:"message"`
}

func (s *SlackResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSlackResult() *SlackResult { _ = "STUB: not implemented"; return nil }
