package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _schedulingConfiguration struct {
	v *types.SchedulingConfiguration
}

func NewSchedulingConfiguration() *_schedulingConfiguration { _ = "STUB: not implemented"; return nil }

func (s *_schedulingConfiguration) AccessControl(accesscontrol types.ConnectorSchedulingVariant) *_schedulingConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_schedulingConfiguration) Full(full types.ConnectorSchedulingVariant) *_schedulingConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_schedulingConfiguration) Incremental(incremental types.ConnectorSchedulingVariant) *_schedulingConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_schedulingConfiguration) SchedulingConfigurationCaster() *types.SchedulingConfiguration {
	_ = "STUB: not implemented"
	return nil
}
