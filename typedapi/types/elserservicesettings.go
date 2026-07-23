package types

type ElserServiceSettings struct {
	AdaptiveAllocations *AdaptiveAllocations `json:"adaptive_allocations,omitempty"`

	NumAllocations int `json:"num_allocations"`

	NumThreads int `json:"num_threads"`
}

func (s *ElserServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewElserServiceSettings() *ElserServiceSettings { _ = "STUB: not implemented"; return nil }

type ElserServiceSettingsVariant interface {
	ElserServiceSettingsCaster() *ElserServiceSettings
}

func (s *ElserServiceSettings) ElserServiceSettingsCaster() *ElserServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
