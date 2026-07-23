package actionstatusoptions

type ActionStatusOptions struct {
	Name string
}

var (
	Success = ActionStatusOptions{"success"}

	Failure = ActionStatusOptions{"failure"}

	Simulated = ActionStatusOptions{"simulated"}

	Throttled = ActionStatusOptions{"throttled"}
)

func (a ActionStatusOptions) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ActionStatusOptions) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a ActionStatusOptions) String() string { _ = "STUB: not implemented"; return "" }
