package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indicatorhealthstatus"
)

type ShardsCapacityIndicator struct {
	Details   *ShardsCapacityIndicatorDetails             `json:"details,omitempty"`
	Diagnosis []Diagnosis                                 `json:"diagnosis,omitempty"`
	Impacts   []Impact                                    `json:"impacts,omitempty"`
	Status    indicatorhealthstatus.IndicatorHealthStatus `json:"status"`
	Symptom   string                                      `json:"symptom"`
}

func (s *ShardsCapacityIndicator) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewShardsCapacityIndicator() *ShardsCapacityIndicator { _ = "STUB: not implemented"; return nil }
