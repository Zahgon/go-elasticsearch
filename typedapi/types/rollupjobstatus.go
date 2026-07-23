package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexingjobstate"
)

type RollupJobStatus struct {
	CurrentPosition map[string]json.RawMessage        `json:"current_position,omitempty"`
	JobState        indexingjobstate.IndexingJobState `json:"job_state"`
	UpgradedDocId   *bool                             `json:"upgraded_doc_id,omitempty"`
}

func (s *RollupJobStatus) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRollupJobStatus() *RollupJobStatus { _ = "STUB: not implemented"; return nil }
