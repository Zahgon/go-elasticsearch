package types

type TrainedModelsRecord struct {
	CreateTime DateTime `json:"create_time,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`

	DataFrameAnalysis *string `json:"data_frame.analysis,omitempty"`

	DataFrameCreateTime *string `json:"data_frame.create_time,omitempty"`

	DataFrameId *string `json:"data_frame.id,omitempty"`

	DataFrameSourceIndex *string `json:"data_frame.source_index,omitempty"`

	Description *string `json:"description,omitempty"`

	HeapSize ByteSize `json:"heap_size,omitempty"`

	Id *string `json:"id,omitempty"`

	IngestCount *string `json:"ingest.count,omitempty"`

	IngestCurrent *string `json:"ingest.current,omitempty"`

	IngestFailed *string `json:"ingest.failed,omitempty"`

	IngestPipelines *string `json:"ingest.pipelines,omitempty"`

	IngestTime *string `json:"ingest.time,omitempty"`

	License *string `json:"license,omitempty"`

	Operations *string `json:"operations,omitempty"`
	Type       *string `json:"type,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (s *TrainedModelsRecord) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelsRecord() *TrainedModelsRecord { _ = "STUB: not implemented"; return nil }
