package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectorstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncstatus"
)

type Connector struct {
	ApiKeyId                         *string                         `json:"api_key_id,omitempty"`
	ApiKeySecretId                   *string                         `json:"api_key_secret_id,omitempty"`
	Configuration                    ConnectorConfiguration          `json:"configuration"`
	CustomScheduling                 ConnectorCustomScheduling       `json:"custom_scheduling"`
	Deleted                          bool                            `json:"deleted"`
	Description                      *string                         `json:"description,omitempty"`
	Error                            *string                         `json:"error,omitempty"`
	Features                         *ConnectorFeatures              `json:"features,omitempty"`
	Filtering                        []FilteringConfig               `json:"filtering"`
	Id                               *string                         `json:"id,omitempty"`
	IndexName                        *string                         `json:"index_name,omitempty"`
	IsNative                         bool                            `json:"is_native"`
	Language                         *string                         `json:"language,omitempty"`
	LastAccessControlSyncError       *string                         `json:"last_access_control_sync_error,omitempty"`
	LastAccessControlSyncScheduledAt DateTime                        `json:"last_access_control_sync_scheduled_at,omitempty"`
	LastAccessControlSyncStatus      *syncstatus.SyncStatus          `json:"last_access_control_sync_status,omitempty"`
	LastDeletedDocumentCount         *int64                          `json:"last_deleted_document_count,omitempty"`
	LastIncrementalSyncScheduledAt   DateTime                        `json:"last_incremental_sync_scheduled_at,omitempty"`
	LastIndexedDocumentCount         *int64                          `json:"last_indexed_document_count,omitempty"`
	LastSeen                         DateTime                        `json:"last_seen,omitempty"`
	LastSyncError                    *string                         `json:"last_sync_error,omitempty"`
	LastSyncScheduledAt              DateTime                        `json:"last_sync_scheduled_at,omitempty"`
	LastSyncStatus                   *syncstatus.SyncStatus          `json:"last_sync_status,omitempty"`
	LastSynced                       DateTime                        `json:"last_synced,omitempty"`
	Name                             *string                         `json:"name,omitempty"`
	Pipeline                         *IngestPipelineParams           `json:"pipeline,omitempty"`
	Scheduling                       SchedulingConfiguration         `json:"scheduling"`
	ServiceType                      *string                         `json:"service_type,omitempty"`
	Status                           connectorstatus.ConnectorStatus `json:"status"`
	SyncCursor                       json.RawMessage                 `json:"sync_cursor,omitempty"`
	SyncNow                          bool                            `json:"sync_now"`
}

func (s *Connector) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewConnector() *Connector { _ = "STUB: not implemented"; return nil }
