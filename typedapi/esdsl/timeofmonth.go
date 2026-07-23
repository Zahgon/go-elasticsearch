package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _timeOfMonth struct {
	v *types.TimeOfMonth
}

func NewTimeOfMonth() *_timeOfMonth { _ = "STUB: not implemented"; return nil }

func (s *_timeOfMonth) At(ats ...string) *_timeOfMonth { _ = "STUB: not implemented"; return nil }

func (s *_timeOfMonth) On(ons ...int) *_timeOfMonth { _ = "STUB: not implemented"; return nil }

func (s *_timeOfMonth) TimeOfMonthCaster() *types.TimeOfMonth {
	_ = "STUB: not implemented"
	return nil
}
