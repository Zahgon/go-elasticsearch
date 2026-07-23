package types

type SchedulingConfiguration struct {
	AccessControl *ConnectorScheduling `json:"access_control,omitempty"`
	Full          *ConnectorScheduling `json:"full,omitempty"`
	Incremental   *ConnectorScheduling `json:"incremental,omitempty"`
}

func NewSchedulingConfiguration() *SchedulingConfiguration { _ = "STUB: not implemented"; return nil }

type SchedulingConfigurationVariant interface {
	SchedulingConfigurationCaster() *SchedulingConfiguration
}

func (s *SchedulingConfiguration) SchedulingConfigurationCaster() *SchedulingConfiguration {
	_ = "STUB: not implemented"
	return nil
}
