package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hourAndMinute struct {
	v *types.HourAndMinute
}

func NewHourAndMinute() *_hourAndMinute { _ = "STUB: not implemented"; return nil }

func (s *_hourAndMinute) Hour(hours ...int) *_hourAndMinute { _ = "STUB: not implemented"; return nil }

func (s *_hourAndMinute) Minute(minutes ...int) *_hourAndMinute {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hourAndMinute) HourAndMinuteCaster() *types.HourAndMinute {
	_ = "STUB: not implemented"
	return nil
}
