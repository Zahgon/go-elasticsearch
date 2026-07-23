package types

type JobTimingStats struct {
	AverageBucketProcessingTimeMs                   Float64 `json:"average_bucket_processing_time_ms,omitempty"`
	BucketCount                                     int64   `json:"bucket_count"`
	ExponentialAverageBucketProcessingTimeMs        Float64 `json:"exponential_average_bucket_processing_time_ms,omitempty"`
	ExponentialAverageBucketProcessingTimePerHourMs Float64 `json:"exponential_average_bucket_processing_time_per_hour_ms"`
	JobId                                           string  `json:"job_id"`
	MaximumBucketProcessingTimeMs                   Float64 `json:"maximum_bucket_processing_time_ms,omitempty"`
	MinimumBucketProcessingTimeMs                   Float64 `json:"minimum_bucket_processing_time_ms,omitempty"`
	TotalBucketProcessingTimeMs                     Float64 `json:"total_bucket_processing_time_ms"`
}

func (s *JobTimingStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJobTimingStats() *JobTimingStats { _ = "STUB: not implemented"; return nil }
