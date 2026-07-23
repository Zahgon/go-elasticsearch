package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/executionstatus"
)

type WatchRecord struct {
	Condition    WatcherCondition                `json:"condition"`
	Input        WatcherInput                    `json:"input"`
	Messages     []string                        `json:"messages"`
	Metadata     Metadata                        `json:"metadata,omitempty"`
	Node         string                          `json:"node"`
	Result       ExecutionResult                 `json:"result"`
	State        executionstatus.ExecutionStatus `json:"state"`
	Status       *WatchStatus                    `json:"status,omitempty"`
	TriggerEvent TriggerEventResult              `json:"trigger_event"`
	User         string                          `json:"user"`
	WatchId      string                          `json:"watch_id"`
}

func (s *WatchRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWatchRecord() *WatchRecord { _ = "STUB: not implemented"; return nil }
