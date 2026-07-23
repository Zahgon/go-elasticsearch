package types

type Influencer struct {
	BucketSpan int64 `json:"bucket_span"`

	Foo *string `json:"foo,omitempty"`

	InfluencerFieldName string `json:"influencer_field_name"`

	InfluencerFieldValue string `json:"influencer_field_value"`

	InfluencerScore Float64 `json:"influencer_score"`

	InitialInfluencerScore Float64 `json:"initial_influencer_score"`

	IsInterim bool `json:"is_interim"`

	JobId string `json:"job_id"`

	Probability Float64 `json:"probability"`

	ResultType string `json:"result_type"`

	Timestamp int64 `json:"timestamp"`
}

func (s *Influencer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInfluencer() *Influencer { _ = "STUB: not implemented"; return nil }
