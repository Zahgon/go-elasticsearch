package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/day"
)

type TimeOfWeek struct {
	At []string  `json:"at"`
	On []day.Day `json:"on"`
}

func NewTimeOfWeek() *TimeOfWeek { _ = "STUB: not implemented"; return nil }

type TimeOfWeekVariant interface {
	TimeOfWeekCaster() *TimeOfWeek
}

func (s *TimeOfWeek) TimeOfWeekCaster() *TimeOfWeek { _ = "STUB: not implemented"; return nil }
