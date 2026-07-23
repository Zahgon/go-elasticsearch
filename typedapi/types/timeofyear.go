package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/month"
)

type TimeOfYear struct {
	At  []string      `json:"at"`
	Int []month.Month `json:"int"`
	On  []int         `json:"on"`
}

func NewTimeOfYear() *TimeOfYear { _ = "STUB: not implemented"; return nil }

type TimeOfYearVariant interface {
	TimeOfYearCaster() *TimeOfYear
}

func (s *TimeOfYear) TimeOfYearCaster() *TimeOfYear { _ = "STUB: not implemented"; return nil }
