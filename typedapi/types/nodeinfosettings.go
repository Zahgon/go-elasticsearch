package types

type NodeInfoSettings struct {
	Action       *NodeInfoAction           `json:"action,omitempty"`
	Bootstrap    *NodeInfoBootstrap        `json:"bootstrap,omitempty"`
	Client       *NodeInfoClient           `json:"client,omitempty"`
	Cluster      NodeInfoSettingsCluster   `json:"cluster"`
	Discovery    *NodeInfoDiscover         `json:"discovery,omitempty"`
	Http         NodeInfoSettingsHttp      `json:"http"`
	Ingest       *NodeInfoSettingsIngest   `json:"ingest,omitempty"`
	Network      *NodeInfoSettingsNetwork  `json:"network,omitempty"`
	Node         NodeInfoSettingsNode      `json:"node"`
	Path         *NodeInfoPath             `json:"path,omitempty"`
	Repositories *NodeInfoRepositories     `json:"repositories,omitempty"`
	Script       *NodeInfoScript           `json:"script,omitempty"`
	Search       *NodeInfoSearch           `json:"search,omitempty"`
	Transport    NodeInfoSettingsTransport `json:"transport"`
	Xpack        *NodeInfoXpack            `json:"xpack,omitempty"`
}

func NewNodeInfoSettings() *NodeInfoSettings { _ = "STUB: not implemented"; return nil }
