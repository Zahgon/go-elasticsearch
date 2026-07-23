package types

type ClusterNodes struct {
	Count ClusterNodeCount `json:"count"`

	DiscoveryTypes map[string]int `json:"discovery_types"`

	Fs               ClusterFileSystem       `json:"fs"`
	IndexingPressure ClusterIndexingPressure `json:"indexing_pressure"`
	Ingest           ClusterIngest           `json:"ingest"`

	Jvm ClusterJvm `json:"jvm"`

	NetworkTypes ClusterNetworkTypes `json:"network_types"`

	Os ClusterOperatingSystem `json:"os"`

	PackagingTypes []NodePackagingType `json:"packaging_types"`

	Plugins []PluginStats `json:"plugins"`

	Process ClusterProcess `json:"process"`

	Versions []string `json:"versions"`
}

func NewClusterNodes() *ClusterNodes { _ = "STUB: not implemented"; return nil }
