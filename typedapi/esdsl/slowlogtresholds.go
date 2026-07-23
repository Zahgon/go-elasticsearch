package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slowlogTresholds struct {
	v *types.SlowlogTresholds
}

func NewSlowlogTresholds() *_slowlogTresholds { _ = "STUB: not implemented"; return nil }

func (s *_slowlogTresholds) Fetch(fetch types.SlowlogTresholdLevelsVariant) *_slowlogTresholds {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slowlogTresholds) Query(query types.SlowlogTresholdLevelsVariant) *_slowlogTresholds {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slowlogTresholds) SlowlogTresholdsCaster() *types.SlowlogTresholds {
	_ = "STUB: not implemented"
	return nil
}
