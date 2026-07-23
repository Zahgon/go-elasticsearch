package types

type AdaptiveAllocationsSettings struct {
	Enabled bool `json:"enabled"`

	MaxNumberOfAllocations *int `json:"max_number_of_allocations,omitempty"`

	MinNumberOfAllocations *int `json:"min_number_of_allocations,omitempty"`
}

func (s *AdaptiveAllocationsSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAdaptiveAllocationsSettings() *AdaptiveAllocationsSettings {
	_ = "STUB: not implemented"
	return nil
}

type AdaptiveAllocationsSettingsVariant interface {
	AdaptiveAllocationsSettingsCaster() *AdaptiveAllocationsSettings
}

func (s *AdaptiveAllocationsSettings) AdaptiveAllocationsSettingsCaster() *AdaptiveAllocationsSettings {
	_ = "STUB: not implemented"
	return nil
}
