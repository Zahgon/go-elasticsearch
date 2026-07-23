package types

type DfsStatisticsBreakdown struct {
	CollectionStatistics      int64 `json:"collection_statistics"`
	CollectionStatisticsCount int64 `json:"collection_statistics_count"`
	CreateWeight              int64 `json:"create_weight"`
	CreateWeightCount         int64 `json:"create_weight_count"`
	Rewrite                   int64 `json:"rewrite"`
	RewriteCount              int64 `json:"rewrite_count"`
	TermStatistics            int64 `json:"term_statistics"`
	TermStatisticsCount       int64 `json:"term_statistics_count"`
}

func (s *DfsStatisticsBreakdown) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDfsStatisticsBreakdown() *DfsStatisticsBreakdown { _ = "STUB: not implemented"; return nil }
