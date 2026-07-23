package syncjobget

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncjobtriggermethod"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncjobtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncstatus"
)

type Response struct {
	CancelationRequestedAt types.DateTime                            `json:"cancelation_requested_at,omitempty"`
	CanceledAt             types.DateTime                            `json:"canceled_at,omitempty"`
	CompletedAt            types.DateTime                            `json:"completed_at,omitempty"`
	Connector              types.SyncJobConnectorReference           `json:"connector"`
	CreatedAt              types.DateTime                            `json:"created_at"`
	DeletedDocumentCount   int64                                     `json:"deleted_document_count"`
	Error                  *string                                   `json:"error,omitempty"`
	Id                     string                                    `json:"id"`
	IndexedDocumentCount   int64                                     `json:"indexed_document_count"`
	IndexedDocumentVolume  int64                                     `json:"indexed_document_volume"`
	JobType                syncjobtype.SyncJobType                   `json:"job_type"`
	LastSeen               types.DateTime                            `json:"last_seen,omitempty"`
	Metadata               map[string]json.RawMessage                `json:"metadata"`
	StartedAt              types.DateTime                            `json:"started_at,omitempty"`
	Status                 syncstatus.SyncStatus                     `json:"status"`
	TotalDocumentCount     int64                                     `json:"total_document_count"`
	TriggerMethod          syncjobtriggermethod.SyncJobTriggerMethod `json:"trigger_method"`
	WorkerHostname         *string                                   `json:"worker_hostname,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
