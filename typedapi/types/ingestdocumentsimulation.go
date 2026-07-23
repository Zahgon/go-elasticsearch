package types

import (
	"encoding/json"
)

type IngestDocumentSimulation struct {
	EffectiveMapping *TypeMapping `json:"effective_mapping,omitempty"`

	Error *ErrorCause `json:"error,omitempty"`

	ExecutedPipelines []string `json:"executed_pipelines"`

	Id_ string `json:"_id"`

	IgnoredFields []map[string]string `json:"ignored_fields,omitempty"`

	Index_                   string            `json:"_index"`
	IngestDocumentSimulation map[string]string `json:"-"`

	Source_  map[string]json.RawMessage `json:"_source"`
	Version_ StringifiedVersionNumber   `json:"_version"`
}

func (s *IngestDocumentSimulation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s IngestDocumentSimulation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIngestDocumentSimulation() *IngestDocumentSimulation { _ = "STUB: not implemented"; return nil }
