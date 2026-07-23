package types

type BucketInfluencer struct {
	AnomalyScore Float64 `json:"anomaly_score"`

	BucketSpan int64 `json:"bucket_span"`

	InfluencerFieldName string `json:"influencer_field_name"`

	InitialAnomalyScore Float64 `json:"initial_anomaly_score"`

	IsInterim bool `json:"is_interim"`

	JobId string `json:"job_id"`

	Probability Float64 `json:"probability"`

	RawAnomalyScore Float64 `json:"raw_anomaly_score"`

	ResultType string `json:"result_type"`

	Timestamp int64 `json:"timestamp"`

	TimestampString DateTime `json:"timestamp_string,omitempty"`
}

func (s *BucketInfluencer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBucketInfluencer() *BucketInfluencer { _ = "STUB: not implemented"; return nil }
