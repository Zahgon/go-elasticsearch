package deploymentassignmentstate

type DeploymentAssignmentState struct {
	Name string
}

var (
	Started = DeploymentAssignmentState{"started"}

	Starting = DeploymentAssignmentState{"starting"}

	Stopping = DeploymentAssignmentState{"stopping"}

	Failed = DeploymentAssignmentState{"failed"}
)

func (d DeploymentAssignmentState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DeploymentAssignmentState) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DeploymentAssignmentState) String() string { _ = "STUB: not implemented"; return "" }
