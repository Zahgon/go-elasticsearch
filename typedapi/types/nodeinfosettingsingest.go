package types

type NodeInfoSettingsIngest struct {
	Append          *NodeInfoIngestInfo `json:"append,omitempty"`
	Attachment      *NodeInfoIngestInfo `json:"attachment,omitempty"`
	Bytes           *NodeInfoIngestInfo `json:"bytes,omitempty"`
	Circle          *NodeInfoIngestInfo `json:"circle,omitempty"`
	Convert         *NodeInfoIngestInfo `json:"convert,omitempty"`
	Csv             *NodeInfoIngestInfo `json:"csv,omitempty"`
	Date            *NodeInfoIngestInfo `json:"date,omitempty"`
	DateIndexName   *NodeInfoIngestInfo `json:"date_index_name,omitempty"`
	Dissect         *NodeInfoIngestInfo `json:"dissect,omitempty"`
	DotExpander     *NodeInfoIngestInfo `json:"dot_expander,omitempty"`
	Drop            *NodeInfoIngestInfo `json:"drop,omitempty"`
	Enrich          *NodeInfoIngestInfo `json:"enrich,omitempty"`
	Fail            *NodeInfoIngestInfo `json:"fail,omitempty"`
	Foreach         *NodeInfoIngestInfo `json:"foreach,omitempty"`
	Geoip           *NodeInfoIngestInfo `json:"geoip,omitempty"`
	Grok            *NodeInfoIngestInfo `json:"grok,omitempty"`
	Gsub            *NodeInfoIngestInfo `json:"gsub,omitempty"`
	Inference       *NodeInfoIngestInfo `json:"inference,omitempty"`
	Join            *NodeInfoIngestInfo `json:"join,omitempty"`
	Json            *NodeInfoIngestInfo `json:"json,omitempty"`
	Kv              *NodeInfoIngestInfo `json:"kv,omitempty"`
	Lowercase       *NodeInfoIngestInfo `json:"lowercase,omitempty"`
	Pipeline        *NodeInfoIngestInfo `json:"pipeline,omitempty"`
	Remove          *NodeInfoIngestInfo `json:"remove,omitempty"`
	Rename          *NodeInfoIngestInfo `json:"rename,omitempty"`
	Script          *NodeInfoIngestInfo `json:"script,omitempty"`
	Set             *NodeInfoIngestInfo `json:"set,omitempty"`
	SetSecurityUser *NodeInfoIngestInfo `json:"set_security_user,omitempty"`
	Sort            *NodeInfoIngestInfo `json:"sort,omitempty"`
	Split           *NodeInfoIngestInfo `json:"split,omitempty"`
	Trim            *NodeInfoIngestInfo `json:"trim,omitempty"`
	Uppercase       *NodeInfoIngestInfo `json:"uppercase,omitempty"`
	Urldecode       *NodeInfoIngestInfo `json:"urldecode,omitempty"`
	UserAgent       *NodeInfoIngestInfo `json:"user_agent,omitempty"`
}

func NewNodeInfoSettingsIngest() *NodeInfoSettingsIngest { _ = "STUB: not implemented"; return nil }
