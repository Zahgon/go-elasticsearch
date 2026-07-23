package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dataframestate"
)

type DataframeAnalytics struct {
	AnalysisStats *DataframeAnalyticsStatsContainer `json:"analysis_stats,omitempty"`

	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`

	DataCounts DataframeAnalyticsStatsDataCounts `json:"data_counts"`

	Id string `json:"id"`

	MemoryUsage DataframeAnalyticsStatsMemoryUsage `json:"memory_usage"`

	Node *NodeAttributes `json:"node,omitempty"`

	Progress []DataframeAnalyticsStatsProgress `json:"progress"`

	State dataframestate.DataframeState `json:"state"`
}

func (s *DataframeAnalytics) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalytics() *DataframeAnalytics { _ = "STUB: not implemented"; return nil }
