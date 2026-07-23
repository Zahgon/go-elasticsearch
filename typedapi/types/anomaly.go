package types

type Anomaly struct {
	Actual []Float64 `json:"actual,omitempty"`

	AnomalyScoreExplanation *AnomalyExplanation `json:"anomaly_score_explanation,omitempty"`

	BucketSpan int64 `json:"bucket_span"`

	ByFieldName *string `json:"by_field_name,omitempty"`

	ByFieldValue *string `json:"by_field_value,omitempty"`

	Causes []AnomalyCause `json:"causes,omitempty"`

	DetectorIndex int `json:"detector_index"`

	FieldName *string `json:"field_name,omitempty"`

	Function *string `json:"function,omitempty"`

	FunctionDescription *string `json:"function_description,omitempty"`

	GeoResults *GeoResults `json:"geo_results,omitempty"`

	Influencers []Influence `json:"influencers,omitempty"`

	InitialRecordScore Float64 `json:"initial_record_score"`

	IsInterim bool `json:"is_interim"`

	JobId string `json:"job_id"`

	OverFieldName *string `json:"over_field_name,omitempty"`

	OverFieldValue *string `json:"over_field_value,omitempty"`

	PartitionFieldName *string `json:"partition_field_name,omitempty"`

	PartitionFieldValue *string `json:"partition_field_value,omitempty"`

	Probability Float64 `json:"probability"`

	RecordScore Float64 `json:"record_score"`

	ResultType string `json:"result_type"`

	Timestamp int64 `json:"timestamp"`

	Typical []Float64 `json:"typical,omitempty"`
}

func (s *Anomaly) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAnomaly() *Anomaly { _ = "STUB: not implemented"; return nil }
