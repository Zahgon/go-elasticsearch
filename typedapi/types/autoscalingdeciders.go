package types

type AutoscalingDeciders struct {
	CurrentCapacity  AutoscalingCapacity           `json:"current_capacity"`
	CurrentNodes     []AutoscalingNode             `json:"current_nodes"`
	Deciders         map[string]AutoscalingDecider `json:"deciders"`
	RequiredCapacity AutoscalingCapacity           `json:"required_capacity"`
}

func NewAutoscalingDeciders() *AutoscalingDeciders { _ = "STUB: not implemented"; return nil }
