package actiontype

type ActionType struct {
	Name string
}

var (
	Email = ActionType{"email"}

	Webhook = ActionType{"webhook"}

	Index = ActionType{"index"}

	Logging = ActionType{"logging"}

	Slack = ActionType{"slack"}

	Pagerduty = ActionType{"pagerduty"}
)

func (a ActionType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ActionType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (a ActionType) String() string { _ = "STUB: not implemented"; return "" }
