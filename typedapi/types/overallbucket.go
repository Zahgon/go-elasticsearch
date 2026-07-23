package types

type OverallBucket struct {
	BucketSpan int64 `json:"bucket_span"`

	IsInterim bool `json:"is_interim"`

	Jobs []OverallBucketJob `json:"jobs"`

	OverallScore Float64 `json:"overall_score"`

	ResultType string `json:"result_type"`

	Timestamp int64 `json:"timestamp"`

	TimestampString DateTime `json:"timestamp_string,omitempty"`
}

func (s *OverallBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewOverallBucket() *OverallBucket { _ = "STUB: not implemented"; return nil }
