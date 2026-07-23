package executionphase

type ExecutionPhase struct {
	Name string
}

var (
	Awaitsexecution = ExecutionPhase{"awaits_execution"}

	Started = ExecutionPhase{"started"}

	Input = ExecutionPhase{"input"}

	Condition = ExecutionPhase{"condition"}

	Actions = ExecutionPhase{"actions"}

	Watchtransform = ExecutionPhase{"watch_transform"}

	Aborted = ExecutionPhase{"aborted"}

	Finished = ExecutionPhase{"finished"}
)

func (e ExecutionPhase) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ExecutionPhase) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e ExecutionPhase) String() string { _ = "STUB: not implemented"; return "" }
