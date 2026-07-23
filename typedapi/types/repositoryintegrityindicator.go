package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indicatorhealthstatus"
)

type RepositoryIntegrityIndicator struct {
	Details   *RepositoryIntegrityIndicatorDetails        `json:"details,omitempty"`
	Diagnosis []Diagnosis                                 `json:"diagnosis,omitempty"`
	Impacts   []Impact                                    `json:"impacts,omitempty"`
	Status    indicatorhealthstatus.IndicatorHealthStatus `json:"status"`
	Symptom   string                                      `json:"symptom"`
}

func (s *RepositoryIntegrityIndicator) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRepositoryIntegrityIndicator() *RepositoryIntegrityIndicator {
	_ = "STUB: not implemented"
	return nil
}
