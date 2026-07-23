package getstats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	PolicyStats                   []types.SnapshotPolicyStats `json:"policy_stats"`
	RetentionDeletionTime         types.Duration              `json:"retention_deletion_time"`
	RetentionDeletionTimeMillis   int64                       `json:"retention_deletion_time_millis"`
	RetentionFailed               int64                       `json:"retention_failed"`
	RetentionRuns                 int64                       `json:"retention_runs"`
	RetentionTimedOut             int64                       `json:"retention_timed_out"`
	TotalSnapshotDeletionFailures int64                       `json:"total_snapshot_deletion_failures"`
	TotalSnapshotsDeleted         int64                       `json:"total_snapshots_deleted"`
	TotalSnapshotsFailed          int64                       `json:"total_snapshots_failed"`
	TotalSnapshotsTaken           int64                       `json:"total_snapshots_taken"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
