package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shutdownstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shutdowntype"
)

type NodeShutdownStatus struct {
	NodeId                string                        `json:"node_id"`
	PersistentTasks       PersistentTaskStatus          `json:"persistent_tasks"`
	Plugins               PluginsStatus                 `json:"plugins"`
	Reason                string                        `json:"reason"`
	ShardMigration        ShardMigrationStatus          `json:"shard_migration"`
	ShutdownStartedmillis int64                         `json:"shutdown_startedmillis"`
	Status                shutdownstatus.ShutdownStatus `json:"status"`
	Type                  shutdowntype.ShutdownType     `json:"type"`
}

func (s *NodeShutdownStatus) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeShutdownStatus() *NodeShutdownStatus { _ = "STUB: not implemented"; return nil }
