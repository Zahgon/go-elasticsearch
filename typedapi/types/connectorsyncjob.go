package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncjobtriggermethod"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncjobtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncstatus"
)

type ConnectorSyncJob struct {
	CancelationRequestedAt DateTime                                  `json:"cancelation_requested_at,omitempty"`
	CanceledAt             DateTime                                  `json:"canceled_at,omitempty"`
	CompletedAt            DateTime                                  `json:"completed_at,omitempty"`
	Connector              SyncJobConnectorReference                 `json:"connector"`
	CreatedAt              DateTime                                  `json:"created_at"`
	DeletedDocumentCount   int64                                     `json:"deleted_document_count"`
	Error                  *string                                   `json:"error,omitempty"`
	Id                     string                                    `json:"id"`
	IndexedDocumentCount   int64                                     `json:"indexed_document_count"`
	IndexedDocumentVolume  int64                                     `json:"indexed_document_volume"`
	JobType                syncjobtype.SyncJobType                   `json:"job_type"`
	LastSeen               DateTime                                  `json:"last_seen,omitempty"`
	Metadata               map[string]json.RawMessage                `json:"metadata"`
	StartedAt              DateTime                                  `json:"started_at,omitempty"`
	Status                 syncstatus.SyncStatus                     `json:"status"`
	TotalDocumentCount     int64                                     `json:"total_document_count"`
	TriggerMethod          syncjobtriggermethod.SyncJobTriggerMethod `json:"trigger_method"`
	WorkerHostname         *string                                   `json:"worker_hostname,omitempty"`
}

func (s *ConnectorSyncJob) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewConnectorSyncJob() *ConnectorSyncJob { _ = "STUB: not implemented"; return nil }
