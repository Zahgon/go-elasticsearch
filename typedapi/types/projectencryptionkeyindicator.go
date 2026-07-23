package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indicatorhealthstatus"
)

type ProjectEncryptionKeyIndicator struct {
	Details   *ProjectEncryptionKeyDetails                `json:"details,omitempty"`
	Diagnosis []Diagnosis                                 `json:"diagnosis,omitempty"`
	Impacts   []Impact                                    `json:"impacts,omitempty"`
	Status    indicatorhealthstatus.IndicatorHealthStatus `json:"status"`
	Symptom   string                                      `json:"symptom"`
}

func (s *ProjectEncryptionKeyIndicator) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewProjectEncryptionKeyIndicator() *ProjectEncryptionKeyIndicator {
	_ = "STUB: not implemented"
	return nil
}
