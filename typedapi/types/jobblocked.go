package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jobblockedreason"
)

type JobBlocked struct {
	Reason jobblockedreason.JobBlockedReason `json:"reason"`
	TaskId *string                           `json:"task_id,omitempty"`
}

func (s *JobBlocked) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJobBlocked() *JobBlocked { _ = "STUB: not implemented"; return nil }
