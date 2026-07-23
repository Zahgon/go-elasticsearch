package types

type TransformIndexerStats struct {
	DeleteTimeInMs                     *int64  `json:"delete_time_in_ms,omitempty"`
	DocumentsDeleted                   *int64  `json:"documents_deleted,omitempty"`
	DocumentsIndexed                   int64   `json:"documents_indexed"`
	DocumentsProcessed                 int64   `json:"documents_processed"`
	ExponentialAvgCheckpointDurationMs Float64 `json:"exponential_avg_checkpoint_duration_ms"`
	ExponentialAvgDocumentsIndexed     Float64 `json:"exponential_avg_documents_indexed"`
	ExponentialAvgDocumentsProcessed   Float64 `json:"exponential_avg_documents_processed"`
	IndexFailures                      int64   `json:"index_failures"`
	IndexTimeInMs                      int64   `json:"index_time_in_ms"`
	IndexTotal                         int64   `json:"index_total"`
	PagesProcessed                     int64   `json:"pages_processed"`
	ProcessingTimeInMs                 int64   `json:"processing_time_in_ms"`
	ProcessingTotal                    int64   `json:"processing_total"`
	SearchFailures                     int64   `json:"search_failures"`
	SearchTimeInMs                     int64   `json:"search_time_in_ms"`
	SearchTotal                        int64   `json:"search_total"`
	TriggerCount                       int64   `json:"trigger_count"`
}

func (s *TransformIndexerStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTransformIndexerStats() *TransformIndexerStats { _ = "STUB: not implemented"; return nil }
