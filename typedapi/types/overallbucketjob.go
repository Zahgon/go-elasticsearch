package types

type OverallBucketJob struct {
	JobId           string  `json:"job_id"`
	MaxAnomalyScore Float64 `json:"max_anomaly_score"`
}

func (s *OverallBucketJob) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewOverallBucketJob() *OverallBucketJob { _ = "STUB: not implemented"; return nil }
