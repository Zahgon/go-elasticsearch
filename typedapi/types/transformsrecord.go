package types

type TransformsRecord struct {
	ChangesLastDetectionTime *string `json:"changes_last_detection_time,omitempty"`

	Checkpoint *string `json:"checkpoint,omitempty"`

	CheckpointDurationTimeExpAvg *string `json:"checkpoint_duration_time_exp_avg,omitempty"`

	CheckpointProgress *string `json:"checkpoint_progress,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	DeleteTime *string `json:"delete_time,omitempty"`

	Description *string `json:"description,omitempty"`

	DestIndex *string `json:"dest_index,omitempty"`

	DocsPerSecond *string `json:"docs_per_second,omitempty"`

	DocumentsDeleted *string `json:"documents_deleted,omitempty"`

	DocumentsIndexed *string `json:"documents_indexed,omitempty"`

	DocumentsProcessed *string `json:"documents_processed,omitempty"`

	Frequency *string `json:"frequency,omitempty"`

	Id *string `json:"id,omitempty"`

	IndexFailure *string `json:"index_failure,omitempty"`

	IndexTime *string `json:"index_time,omitempty"`

	IndexTotal *string `json:"index_total,omitempty"`

	IndexedDocumentsExpAvg *string `json:"indexed_documents_exp_avg,omitempty"`

	LastSearchTime *string `json:"last_search_time,omitempty"`

	MaxPageSearchSize *string `json:"max_page_search_size,omitempty"`

	PagesProcessed *string `json:"pages_processed,omitempty"`

	Pipeline *string `json:"pipeline,omitempty"`

	ProcessedDocumentsExpAvg *string `json:"processed_documents_exp_avg,omitempty"`

	ProcessingTime *string `json:"processing_time,omitempty"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	Reason *string `json:"reason,omitempty"`

	SearchFailure *string `json:"search_failure,omitempty"`

	SearchTime *string `json:"search_time,omitempty"`

	SearchTotal *string `json:"search_total,omitempty"`

	SourceIndex *string `json:"source_index,omitempty"`

	State *string `json:"state,omitempty"`

	TransformType *string `json:"transform_type,omitempty"`

	TriggerCount *string `json:"trigger_count,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (s *TransformsRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTransformsRecord() *TransformsRecord { _ = "STUB: not implemented"; return nil }
