package types

type AggregationBreakdown struct {
	BuildAggregation        int64  `json:"build_aggregation"`
	BuildAggregationCount   int64  `json:"build_aggregation_count"`
	BuildLeafCollector      int64  `json:"build_leaf_collector"`
	BuildLeafCollectorCount int64  `json:"build_leaf_collector_count"`
	Collect                 int64  `json:"collect"`
	CollectCount            int64  `json:"collect_count"`
	Initialize              int64  `json:"initialize"`
	InitializeCount         int64  `json:"initialize_count"`
	PostCollection          *int64 `json:"post_collection,omitempty"`
	PostCollectionCount     *int64 `json:"post_collection_count,omitempty"`
	Reduce                  int64  `json:"reduce"`
	ReduceCount             int64  `json:"reduce_count"`
}

func (s *AggregationBreakdown) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAggregationBreakdown() *AggregationBreakdown { _ = "STUB: not implemented"; return nil }
