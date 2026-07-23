package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/retentionsource"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/samplingmethod"
)

type _dataStreamLifecycle struct {
	v *types.DataStreamLifecycle
}

func NewDataStreamLifecycle() *_dataStreamLifecycle { _ = "STUB: not implemented"; return nil }

func (s *_dataStreamLifecycle) DataRetention(duration types.DurationVariant) *_dataStreamLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamLifecycle) Downsampling(downsamplings ...types.DownsamplingRoundVariant) *_dataStreamLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamLifecycle) DownsamplingValues(downsamplingvalues []types.DownsamplingRound) *_dataStreamLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamLifecycle) DownsamplingMethod(downsamplingmethod samplingmethod.SamplingMethod) *_dataStreamLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamLifecycle) EffectiveRetention(duration types.DurationVariant) *_dataStreamLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamLifecycle) Enabled(enabled bool) *_dataStreamLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamLifecycle) FrozenAfter(duration types.DurationVariant) *_dataStreamLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamLifecycle) RetentionDeterminedBy(retentiondeterminedby retentionsource.RetentionSource) *_dataStreamLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamLifecycle) DataStreamLifecycleCaster() *types.DataStreamLifecycle {
	_ = "STUB: not implemented"
	return nil
}
