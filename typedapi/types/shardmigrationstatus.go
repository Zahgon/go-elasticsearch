package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shutdownstatus"
)

type ShardMigrationStatus struct {
	Status shutdownstatus.ShutdownStatus `json:"status"`
}

func NewShardMigrationStatus() *ShardMigrationStatus { _ = "STUB: not implemented"; return nil }
