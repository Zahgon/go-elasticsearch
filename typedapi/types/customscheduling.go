package types

type CustomScheduling struct {
	ConfigurationOverrides CustomSchedulingConfigurationOverrides `json:"configuration_overrides"`
	Enabled                bool                                   `json:"enabled"`
	Interval               string                                 `json:"interval"`
	LastSynced             DateTime                               `json:"last_synced,omitempty"`
	Name                   string                                 `json:"name"`
}

func (s *CustomScheduling) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCustomScheduling() *CustomScheduling { _ = "STUB: not implemented"; return nil }
