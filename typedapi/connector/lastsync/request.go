package lastsync

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncstatus"
)

type Request struct {
	LastAccessControlSyncError       *string                `json:"last_access_control_sync_error,omitempty"`
	LastAccessControlSyncScheduledAt types.DateTime         `json:"last_access_control_sync_scheduled_at,omitempty"`
	LastAccessControlSyncStatus      *syncstatus.SyncStatus `json:"last_access_control_sync_status,omitempty"`
	LastDeletedDocumentCount         *int64                 `json:"last_deleted_document_count,omitempty"`
	LastIncrementalSyncScheduledAt   types.DateTime         `json:"last_incremental_sync_scheduled_at,omitempty"`
	LastIndexedDocumentCount         *int64                 `json:"last_indexed_document_count,omitempty"`
	LastSeen                         types.DateTime         `json:"last_seen,omitempty"`
	LastSyncError                    *string                `json:"last_sync_error,omitempty"`
	LastSyncScheduledAt              types.DateTime         `json:"last_sync_scheduled_at,omitempty"`
	LastSyncStatus                   *syncstatus.SyncStatus `json:"last_sync_status,omitempty"`
	LastSynced                       types.DateTime         `json:"last_synced,omitempty"`
	SyncCursor                       json.RawMessage        `json:"sync_cursor,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
