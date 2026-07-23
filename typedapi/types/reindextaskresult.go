package types

type ReindexTaskResult struct {
	Batches *int64 `json:"batches,omitempty"`

	Created *int64 `json:"created,omitempty"`

	Deleted *int64 `json:"deleted,omitempty"`

	Failures []BulkIndexByScrollFailure `json:"failures,omitempty"`

	Noops *int64 `json:"noops,omitempty"`

	RequestsPerSecond *float32 `json:"requests_per_second,omitempty"`

	Retries *Retries `json:"retries,omitempty"`

	ThrottledMillis *int64 `json:"throttled_millis,omitempty"`

	ThrottledUntilMillis *int64 `json:"throttled_until_millis,omitempty"`

	TimedOut *bool `json:"timed_out,omitempty"`

	Took *int64 `json:"took,omitempty"`

	Total *int64 `json:"total,omitempty"`

	Updated *int64 `json:"updated,omitempty"`

	VersionConflicts *int64 `json:"version_conflicts,omitempty"`
}

func (s *ReindexTaskResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReindexTaskResult() *ReindexTaskResult { _ = "STUB: not implemented"; return nil }
