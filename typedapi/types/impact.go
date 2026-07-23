package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/impactarea"
)

type Impact struct {
	Description string                  `json:"description"`
	Id          string                  `json:"id"`
	ImpactAreas []impactarea.ImpactArea `json:"impact_areas"`
	Severity    int                     `json:"severity"`
}

func (s *Impact) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewImpact() *Impact { _ = "STUB: not implemented"; return nil }
