package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shardsstatsstage"
)

type SnapshotShardsStatus struct {
	Stage shardsstatsstage.ShardsStatsStage `json:"stage"`
	Stats ShardsStatsSummary                `json:"stats"`
}

func NewSnapshotShardsStatus() *SnapshotShardsStatus { _ = "STUB: not implemented"; return nil }
