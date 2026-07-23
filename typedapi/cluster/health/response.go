package health

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
)

type Response struct {
	ActivePrimaryShards int `json:"active_primary_shards"`

	ActiveShards int `json:"active_shards"`

	ActiveShardsPercent *string `json:"active_shards_percent,omitempty"`

	ActiveShardsPercentAsNumber types.Float64 `json:"active_shards_percent_as_number"`

	ClusterName string `json:"cluster_name"`

	DelayedUnassignedShards int                               `json:"delayed_unassigned_shards"`
	Indices                 map[string]types.IndexHealthStats `json:"indices,omitempty"`

	InitializingShards int `json:"initializing_shards"`

	NumberOfDataNodes int `json:"number_of_data_nodes"`

	NumberOfInFlightFetch int `json:"number_of_in_flight_fetch"`

	NumberOfNodes int `json:"number_of_nodes"`

	NumberOfPendingTasks int `json:"number_of_pending_tasks"`

	RelocatingShards int                       `json:"relocating_shards"`
	Status           healthstatus.HealthStatus `json:"status"`

	TaskMaxWaitingInQueue types.Duration `json:"task_max_waiting_in_queue,omitempty"`

	TaskMaxWaitingInQueueMillis int64 `json:"task_max_waiting_in_queue_millis"`

	TimedOut bool `json:"timed_out"`

	UnassignedPrimaryShards int `json:"unassigned_primary_shards"`

	UnassignedShards int `json:"unassigned_shards"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
