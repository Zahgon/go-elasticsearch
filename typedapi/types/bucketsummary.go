package types

type BucketSummary struct {
	AnomalyScore      Float64            `json:"anomaly_score"`
	BucketInfluencers []BucketInfluencer `json:"bucket_influencers"`

	BucketSpan int64 `json:"bucket_span"`

	EventCount int64 `json:"event_count"`

	InitialAnomalyScore Float64 `json:"initial_anomaly_score"`

	IsInterim bool `json:"is_interim"`

	JobId string `json:"job_id"`

	ProcessingTimeMs int64 `json:"processing_time_ms"`

	ResultType string `json:"result_type"`

	Timestamp int64 `json:"timestamp"`

	TimestampString DateTime `json:"timestamp_string,omitempty"`
}

func (s *BucketSummary) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBucketSummary() *BucketSummary { _ = "STUB: not implemented"; return nil }
