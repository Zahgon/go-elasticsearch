package types

type AutoscalingResources struct {
	Memory  int `json:"memory"`
	Storage int `json:"storage"`
}

func (s *AutoscalingResources) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAutoscalingResources() *AutoscalingResources { _ = "STUB: not implemented"; return nil }
