package postdata

type Response struct {
	BucketCount                 int64  `json:"bucket_count"`
	EarliestRecordTimestamp     *int64 `json:"earliest_record_timestamp,omitempty"`
	EmptyBucketCount            int64  `json:"empty_bucket_count"`
	InputBytes                  int64  `json:"input_bytes"`
	InputFieldCount             int64  `json:"input_field_count"`
	InputRecordCount            int64  `json:"input_record_count"`
	InvalidDateCount            int64  `json:"invalid_date_count"`
	JobId                       string `json:"job_id"`
	LastDataTime                *int64 `json:"last_data_time,omitempty"`
	LatestEmptyBucketTimestamp  *int64 `json:"latest_empty_bucket_timestamp,omitempty"`
	LatestRecordTimestamp       *int64 `json:"latest_record_timestamp,omitempty"`
	LatestSparseBucketTimestamp *int64 `json:"latest_sparse_bucket_timestamp,omitempty"`
	LogTime                     *int64 `json:"log_time,omitempty"`
	MissingFieldCount           int64  `json:"missing_field_count"`
	OutOfOrderTimestampCount    int64  `json:"out_of_order_timestamp_count"`
	ProcessedFieldCount         int64  `json:"processed_field_count"`
	ProcessedRecordCount        int64  `json:"processed_record_count"`
	SparseBucketCount           int64  `json:"sparse_bucket_count"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
