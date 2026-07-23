package types

type ParentReindexStatus struct {
	Batches int64 `json:"batches"`

	Cancelled *string `json:"cancelled,omitempty"`

	Created *int64 `json:"created,omitempty"`

	Deleted int64 `json:"deleted"`

	Noops int64 `json:"noops"`

	RequestsPerSecond float32 `json:"requests_per_second"`

	Retries Retries `json:"retries"`

	SliceId   *int            `json:"slice_id,omitempty"`
	Slices    []ReindexStatus `json:"slices,omitempty"`
	Throttled Duration        `json:"throttled,omitempty"`

	ThrottledMillis int64    `json:"throttled_millis"`
	ThrottledUntil  Duration `json:"throttled_until,omitempty"`

	ThrottledUntilMillis int64 `json:"throttled_until_millis"`

	Total int64 `json:"total"`

	Updated *int64 `json:"updated,omitempty"`

	VersionConflicts int64 `json:"version_conflicts"`
}

func (s *ParentReindexStatus) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewParentReindexStatus() *ParentReindexStatus { _ = "STUB: not implemented"; return nil }
