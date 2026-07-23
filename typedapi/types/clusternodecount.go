package types

type ClusterNodeCount struct {
	CoordinatingOnly    *int `json:"coordinating_only,omitempty"`
	Data                *int `json:"data,omitempty"`
	DataCold            *int `json:"data_cold,omitempty"`
	DataContent         *int `json:"data_content,omitempty"`
	DataFrozen          *int `json:"data_frozen,omitempty"`
	DataHot             *int `json:"data_hot,omitempty"`
	DataWarm            *int `json:"data_warm,omitempty"`
	Index               *int `json:"index,omitempty"`
	Ingest              *int `json:"ingest,omitempty"`
	Master              *int `json:"master,omitempty"`
	Ml                  *int `json:"ml,omitempty"`
	RemoteClusterClient *int `json:"remote_cluster_client,omitempty"`
	Search              *int `json:"search,omitempty"`
	Total               int  `json:"total"`
	Transform           *int `json:"transform,omitempty"`
	VotingOnly          *int `json:"voting_only,omitempty"`
}

func (s *ClusterNodeCount) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterNodeCount() *ClusterNodeCount { _ = "STUB: not implemented"; return nil }
