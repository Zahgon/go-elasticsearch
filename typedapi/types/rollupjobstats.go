package types

type RollupJobStats struct {
	DocumentsProcessed int64 `json:"documents_processed"`
	IndexFailures      int64 `json:"index_failures"`
	IndexTimeInMs      int64 `json:"index_time_in_ms"`
	IndexTotal         int64 `json:"index_total"`
	PagesProcessed     int64 `json:"pages_processed"`
	ProcessingTimeInMs int64 `json:"processing_time_in_ms"`
	ProcessingTotal    int64 `json:"processing_total"`
	RollupsIndexed     int64 `json:"rollups_indexed"`
	SearchFailures     int64 `json:"search_failures"`
	SearchTimeInMs     int64 `json:"search_time_in_ms"`
	SearchTotal        int64 `json:"search_total"`
	TriggerCount       int64 `json:"trigger_count"`
}

func (s *RollupJobStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRollupJobStats() *RollupJobStats { _ = "STUB: not implemented"; return nil }
