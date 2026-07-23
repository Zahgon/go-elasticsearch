package types

type DataTierPhaseStatistics struct {
	DocCount                    int64 `json:"doc_count"`
	IndexCount                  int64 `json:"index_count"`
	NodeCount                   int64 `json:"node_count"`
	PrimaryShardCount           int64 `json:"primary_shard_count"`
	PrimaryShardSizeAvgBytes    int64 `json:"primary_shard_size_avg_bytes"`
	PrimaryShardSizeMadBytes    int64 `json:"primary_shard_size_mad_bytes"`
	PrimaryShardSizeMedianBytes int64 `json:"primary_shard_size_median_bytes"`
	PrimarySizeBytes            int64 `json:"primary_size_bytes"`
	TotalShardCount             int64 `json:"total_shard_count"`
	TotalSizeBytes              int64 `json:"total_size_bytes"`
}

func (s *DataTierPhaseStatistics) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataTierPhaseStatistics() *DataTierPhaseStatistics { _ = "STUB: not implemented"; return nil }
