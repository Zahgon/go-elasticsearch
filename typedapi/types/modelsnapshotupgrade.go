package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/snapshotupgradestate"
)

type ModelSnapshotUpgrade struct {
	AssignmentExplanation string                                    `json:"assignment_explanation"`
	JobId                 string                                    `json:"job_id"`
	Node                  DiscoveryNode                             `json:"node"`
	SnapshotId            string                                    `json:"snapshot_id"`
	State                 snapshotupgradestate.SnapshotUpgradeState `json:"state"`
}

func (s *ModelSnapshotUpgrade) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewModelSnapshotUpgrade() *ModelSnapshotUpgrade { _ = "STUB: not implemented"; return nil }
