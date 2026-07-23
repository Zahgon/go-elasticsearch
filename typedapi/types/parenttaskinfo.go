package types

import (
	"encoding/json"
)

type ParentTaskInfo struct {
	Action      string     `json:"action"`
	Cancellable bool       `json:"cancellable"`
	Cancelled   *bool      `json:"cancelled,omitempty"`
	Children    []TaskInfo `json:"children,omitempty"`

	Description *string           `json:"description,omitempty"`
	Headers     map[string]string `json:"headers"`
	Id          int64             `json:"id"`
	Node        string            `json:"node"`

	OriginalStartTime *string `json:"original_start_time,omitempty"`

	OriginalStartTimeInMillis *int64 `json:"original_start_time_in_millis,omitempty"`

	OriginalTaskId     *string  `json:"original_task_id,omitempty"`
	ParentTaskId       *string  `json:"parent_task_id,omitempty"`
	RunningTime        Duration `json:"running_time,omitempty"`
	RunningTimeInNanos int64    `json:"running_time_in_nanos"`
	StartTimeInMillis  int64    `json:"start_time_in_millis"`

	Status json.RawMessage `json:"status,omitempty"`
	Type   string          `json:"type"`
}

func (s *ParentTaskInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewParentTaskInfo() *ParentTaskInfo { _ = "STUB: not implemented"; return nil }
