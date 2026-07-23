package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slowlogTresholdLevels struct {
	v *types.SlowlogTresholdLevels
}

func NewSlowlogTresholdLevels() *_slowlogTresholdLevels { _ = "STUB: not implemented"; return nil }

func (s *_slowlogTresholdLevels) Debug(duration types.DurationVariant) *_slowlogTresholdLevels {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slowlogTresholdLevels) Info(duration types.DurationVariant) *_slowlogTresholdLevels {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slowlogTresholdLevels) Trace(duration types.DurationVariant) *_slowlogTresholdLevels {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slowlogTresholdLevels) Warn(duration types.DurationVariant) *_slowlogTresholdLevels {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slowlogTresholdLevels) SlowlogTresholdLevelsCaster() *types.SlowlogTresholdLevels {
	_ = "STUB: not implemented"
	return nil
}
