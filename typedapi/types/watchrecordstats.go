package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/executionphase"
)

type WatchRecordStats struct {
	ExecutedActions []string `json:"executed_actions,omitempty"`

	ExecutionPhase executionphase.ExecutionPhase `json:"execution_phase"`

	ExecutionTime DateTime `json:"execution_time"`

	TriggeredTime DateTime `json:"triggered_time"`
	WatchId       string   `json:"watch_id"`

	WatchRecordId string `json:"watch_record_id"`
}

func (s *WatchRecordStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWatchRecordStats() *WatchRecordStats { _ = "STUB: not implemented"; return nil }
