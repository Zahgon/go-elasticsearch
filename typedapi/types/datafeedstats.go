package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/datafeedstate"
)

type DatafeedStats struct {
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`

	DatafeedId string `json:"datafeed_id"`

	Node *DiscoveryNodeCompact `json:"node,omitempty"`

	RunningState *DatafeedRunningState `json:"running_state,omitempty"`

	State datafeedstate.DatafeedState `json:"state"`

	TimingStats *DatafeedTimingStats `json:"timing_stats,omitempty"`
}

func (s *DatafeedStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDatafeedStats() *DatafeedStats { _ = "STUB: not implemented"; return nil }
