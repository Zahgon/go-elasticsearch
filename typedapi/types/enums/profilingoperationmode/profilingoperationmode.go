package profilingoperationmode

type ProfilingOperationMode struct {
	Name string
}

var (
	RUNNING = ProfilingOperationMode{"RUNNING"}

	STOPPING = ProfilingOperationMode{"STOPPING"}

	STOPPED = ProfilingOperationMode{"STOPPED"}
)

func (p ProfilingOperationMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ProfilingOperationMode) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (p ProfilingOperationMode) String() string { _ = "STUB: not implemented"; return "" }
