package syncjobpost

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncjobtriggermethod"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncjobtype"
)

type Request struct {
	Id            string                                     `json:"id"`
	JobType       *syncjobtype.SyncJobType                   `json:"job_type,omitempty"`
	TriggerMethod *syncjobtriggermethod.SyncJobTriggerMethod `json:"trigger_method,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
