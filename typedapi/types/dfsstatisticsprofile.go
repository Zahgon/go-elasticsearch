package types

import (
	"encoding/json"
)

type DfsStatisticsProfile struct {
	Breakdown   DfsStatisticsBreakdown     `json:"breakdown"`
	Children    []DfsStatisticsProfile     `json:"children,omitempty"`
	Debug       map[string]json.RawMessage `json:"debug,omitempty"`
	Description string                     `json:"description"`
	Time        Duration                   `json:"time,omitempty"`
	TimeInNanos int64                      `json:"time_in_nanos"`
	Type        string                     `json:"type"`
}

func (s *DfsStatisticsProfile) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDfsStatisticsProfile() *DfsStatisticsProfile { _ = "STUB: not implemented"; return nil }
