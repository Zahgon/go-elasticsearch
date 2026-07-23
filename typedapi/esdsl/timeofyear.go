package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/month"
)

type _timeOfYear struct {
	v *types.TimeOfYear
}

func NewTimeOfYear() *_timeOfYear { _ = "STUB: not implemented"; return nil }

func (s *_timeOfYear) At(ats ...string) *_timeOfYear { _ = "STUB: not implemented"; return nil }

func (s *_timeOfYear) Int(ints ...month.Month) *_timeOfYear { _ = "STUB: not implemented"; return nil }

func (s *_timeOfYear) On(ons ...int) *_timeOfYear { _ = "STUB: not implemented"; return nil }

func (s *_timeOfYear) TimeOfYearCaster() *types.TimeOfYear { _ = "STUB: not implemented"; return nil }
