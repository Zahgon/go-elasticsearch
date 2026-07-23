package pipelinesimulationstatusoptions

type PipelineSimulationStatusOptions struct {
	Name string
}

var (
	Success = PipelineSimulationStatusOptions{"success"}

	Error = PipelineSimulationStatusOptions{"error"}

	Errorignored = PipelineSimulationStatusOptions{"error_ignored"}

	Skipped = PipelineSimulationStatusOptions{"skipped"}

	Dropped = PipelineSimulationStatusOptions{"dropped"}
)

func (p PipelineSimulationStatusOptions) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PipelineSimulationStatusOptions) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PipelineSimulationStatusOptions) String() string { _ = "STUB: not implemented"; return "" }
