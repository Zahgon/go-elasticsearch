package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indicatorhealthstatus"
)

type DiskIndicator struct {
	Details   *DiskIndicatorDetails                       `json:"details,omitempty"`
	Diagnosis []Diagnosis                                 `json:"diagnosis,omitempty"`
	Impacts   []Impact                                    `json:"impacts,omitempty"`
	Status    indicatorhealthstatus.IndicatorHealthStatus `json:"status"`
	Symptom   string                                      `json:"symptom"`
}

func (s *DiskIndicator) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDiskIndicator() *DiskIndicator { _ = "STUB: not implemented"; return nil }
