package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/day"
)

type _timeOfWeek struct {
	v *types.TimeOfWeek
}

func NewTimeOfWeek() *_timeOfWeek { _ = "STUB: not implemented"; return nil }

func (s *_timeOfWeek) At(ats ...string) *_timeOfWeek { _ = "STUB: not implemented"; return nil }

func (s *_timeOfWeek) On(ons ...day.Day) *_timeOfWeek { _ = "STUB: not implemented"; return nil }

func (s *_timeOfWeek) TimeOfWeekCaster() *types.TimeOfWeek { _ = "STUB: not implemented"; return nil }
