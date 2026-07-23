package deploymentallocationstate

type DeploymentAllocationState struct {
	Name string
}

var (
	Started = DeploymentAllocationState{"started"}

	Starting = DeploymentAllocationState{"starting"}

	Fullyallocated = DeploymentAllocationState{"fully_allocated"}
)

func (d DeploymentAllocationState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DeploymentAllocationState) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DeploymentAllocationState) String() string { _ = "STUB: not implemented"; return "" }
