package types

import (
	"encoding/json"
)

type SyncJobConnectorReference struct {
	Configuration ConnectorConfiguration `json:"configuration"`
	Filtering     FilteringRules         `json:"filtering"`
	Id            string                 `json:"id"`
	IndexName     string                 `json:"index_name"`
	Language      *string                `json:"language,omitempty"`
	Pipeline      *IngestPipelineParams  `json:"pipeline,omitempty"`
	ServiceType   string                 `json:"service_type"`
	SyncCursor    json.RawMessage        `json:"sync_cursor,omitempty"`
}

func (s *SyncJobConnectorReference) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSyncJobConnectorReference() *SyncJobConnectorReference {
	_ = "STUB: not implemented"
	return nil
}
