package types

type XpackFeatures struct {
	AggregateMetric     XpackFeature  `json:"aggregate_metric"`
	Analytics           XpackFeature  `json:"analytics"`
	Archive             XpackFeature  `json:"archive"`
	Ccr                 XpackFeature  `json:"ccr"`
	DataStreams         XpackFeature  `json:"data_streams"`
	DataTiers           XpackFeature  `json:"data_tiers"`
	Enrich              XpackFeature  `json:"enrich"`
	EnterpriseSearch    XpackFeature  `json:"enterprise_search"`
	Eql                 XpackFeature  `json:"eql"`
	Esql                XpackFeature  `json:"esql"`
	GpuVectorIndexing   XpackFeature  `json:"gpu_vector_indexing"`
	Graph               XpackFeature  `json:"graph"`
	Ilm                 XpackFeature  `json:"ilm"`
	Logsdb              XpackFeature  `json:"logsdb"`
	Logstash            XpackFeature  `json:"logstash"`
	Ml                  XpackFeature  `json:"ml"`
	Monitoring          XpackFeature  `json:"monitoring"`
	Rollup              XpackFeature  `json:"rollup"`
	RuntimeFields       *XpackFeature `json:"runtime_fields,omitempty"`
	SearchableSnapshots XpackFeature  `json:"searchable_snapshots"`
	Security            XpackFeature  `json:"security"`
	Slm                 XpackFeature  `json:"slm"`
	Spatial             XpackFeature  `json:"spatial"`
	Sql                 XpackFeature  `json:"sql"`
	Transform           XpackFeature  `json:"transform"`
	UniversalProfiling  XpackFeature  `json:"universal_profiling"`
	VotingOnly          XpackFeature  `json:"voting_only"`
	Watcher             XpackFeature  `json:"watcher"`
}

func NewXpackFeatures() *XpackFeatures { _ = "STUB: not implemented"; return nil }
