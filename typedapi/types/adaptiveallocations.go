package types

type AdaptiveAllocations struct {
	Enabled *bool `json:"enabled,omitempty"`

	MaxNumberOfAllocations *int `json:"max_number_of_allocations,omitempty"`

	MinNumberOfAllocations *int `json:"min_number_of_allocations,omitempty"`
}

func (s *AdaptiveAllocations) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAdaptiveAllocations() *AdaptiveAllocations { _ = "STUB: not implemented"; return nil }

type AdaptiveAllocationsVariant interface {
	AdaptiveAllocationsCaster() *AdaptiveAllocations
}

func (s *AdaptiveAllocations) AdaptiveAllocationsCaster() *AdaptiveAllocations {
	_ = "STUB: not implemented"
	return nil
}
