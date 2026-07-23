package types

type IngestStats struct {
	Count int64 `json:"count"`

	Current int64 `json:"current"`

	Failed int64 `json:"failed"`

	IngestedAsFirstPipelineInBytes int64 `json:"ingested_as_first_pipeline_in_bytes"`

	Processors []map[string]KeyedProcessor `json:"processors"`

	ProducedAsFirstPipelineInBytes int64 `json:"produced_as_first_pipeline_in_bytes"`

	TimeInMillis int64 `json:"time_in_millis"`
}

func (s *IngestStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIngestStats() *IngestStats { _ = "STUB: not implemented"; return nil }
