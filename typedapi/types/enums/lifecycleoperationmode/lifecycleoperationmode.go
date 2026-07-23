package lifecycleoperationmode

type LifecycleOperationMode struct {
	Name string
}

var (
	RUNNING = LifecycleOperationMode{"RUNNING"}

	STOPPING = LifecycleOperationMode{"STOPPING"}

	STOPPED = LifecycleOperationMode{"STOPPED"}
)

func (l LifecycleOperationMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *LifecycleOperationMode) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (l LifecycleOperationMode) String() string { _ = "STUB: not implemented"; return "" }
