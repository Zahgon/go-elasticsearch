package actionexecutionmode

type ActionExecutionMode struct {
	Name string
}

var (
	Simulate = ActionExecutionMode{"simulate"}

	Forcesimulate = ActionExecutionMode{"force_simulate"}

	Execute = ActionExecutionMode{"execute"}

	Forceexecute = ActionExecutionMode{"force_execute"}

	Skip = ActionExecutionMode{"skip"}
)

func (a ActionExecutionMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ActionExecutionMode) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a ActionExecutionMode) String() string { _ = "STUB: not implemented"; return "" }
