package types

type Statistics struct {
	Policy                        *string  `json:"policy,omitempty"`
	RetentionDeletionTime         Duration `json:"retention_deletion_time,omitempty"`
	RetentionDeletionTimeMillis   *int64   `json:"retention_deletion_time_millis,omitempty"`
	RetentionFailed               *int64   `json:"retention_failed,omitempty"`
	RetentionRuns                 *int64   `json:"retention_runs,omitempty"`
	RetentionTimedOut             *int64   `json:"retention_timed_out,omitempty"`
	TotalSnapshotDeletionFailures *int64   `json:"total_snapshot_deletion_failures,omitempty"`
	TotalSnapshotsDeleted         *int64   `json:"total_snapshots_deleted,omitempty"`
	TotalSnapshotsFailed          *int64   `json:"total_snapshots_failed,omitempty"`
	TotalSnapshotsTaken           *int64   `json:"total_snapshots_taken,omitempty"`
}

func (s *Statistics) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStatistics() *Statistics { _ = "STUB: not implemented"; return nil }
