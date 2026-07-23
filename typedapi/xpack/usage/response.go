package usage

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	AggregateMetric     types.Base                    `json:"aggregate_metric"`
	Analytics           types.Analytics               `json:"analytics"`
	Archive             types.Archive                 `json:"archive"`
	Ccr                 types.Ccr                     `json:"ccr"`
	DataFrame           *types.Base                   `json:"data_frame,omitempty"`
	DataScience         *types.Base                   `json:"data_science,omitempty"`
	DataStreams         *types.DataStreams            `json:"data_streams,omitempty"`
	DataTiers           types.DataTiers               `json:"data_tiers"`
	Enrich              *types.Base                   `json:"enrich,omitempty"`
	Eql                 types.Eql                     `json:"eql"`
	Flattened           *types.Flattened              `json:"flattened,omitempty"`
	GpuVectorIndexing   *types.GpuVectorIndexing      `json:"gpu_vector_indexing,omitempty"`
	Graph               types.Base                    `json:"graph"`
	HealthApi           *types.HealthStatistics       `json:"health_api,omitempty"`
	Ilm                 types.Ilm                     `json:"ilm"`
	Logstash            types.Base                    `json:"logstash"`
	Ml                  types.MachineLearning         `json:"ml"`
	Monitoring          types.Monitoring              `json:"monitoring"`
	Rollup              types.Base                    `json:"rollup"`
	RuntimeFields       *types.XpackRuntimeFieldTypes `json:"runtime_fields,omitempty"`
	SearchableSnapshots types.SearchableSnapshots     `json:"searchable_snapshots"`
	Security            types.Security                `json:"security"`
	Slm                 types.Slm                     `json:"slm"`
	Spatial             types.Base                    `json:"spatial"`
	Sql                 types.Sql                     `json:"sql"`
	Transform           types.Base                    `json:"transform"`
	Vectors             *types.Vector                 `json:"vectors,omitempty"`
	VotingOnly          types.Base                    `json:"voting_only"`
	Watcher             types.Watcher                 `json:"watcher"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
