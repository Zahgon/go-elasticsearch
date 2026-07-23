package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/retentionsource"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/samplingmethod"
)

type DataStreamLifecycleWithRollover struct {
	DataRetention Duration `json:"data_retention,omitempty"`

	Downsampling []DownsamplingRound `json:"downsampling,omitempty"`

	DownsamplingMethod *samplingmethod.SamplingMethod `json:"downsampling_method,omitempty"`

	EffectiveRetention Duration `json:"effective_retention,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	FrozenAfter Duration `json:"frozen_after,omitempty"`

	RetentionDeterminedBy *retentionsource.RetentionSource `json:"retention_determined_by,omitempty"`

	Rollover *DataStreamLifecycleRolloverConditions `json:"rollover,omitempty"`
}

func (s *DataStreamLifecycleWithRollover) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamLifecycleWithRollover() *DataStreamLifecycleWithRollover {
	_ = "STUB: not implemented"
	return nil
}
