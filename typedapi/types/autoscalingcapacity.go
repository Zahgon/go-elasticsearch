package types

type AutoscalingCapacity struct {
	Node  AutoscalingResources `json:"node"`
	Total AutoscalingResources `json:"total"`
}

func NewAutoscalingCapacity() *AutoscalingCapacity { _ = "STUB: not implemented"; return nil }
