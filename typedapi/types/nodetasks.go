package types

type NodeTasks struct {
	Attributes       map[string]string   `json:"attributes,omitempty"`
	Host             *string             `json:"host,omitempty"`
	Ip               *string             `json:"ip,omitempty"`
	Name             *string             `json:"name,omitempty"`
	Roles            []string            `json:"roles,omitempty"`
	Tasks            map[string]TaskInfo `json:"tasks"`
	TransportAddress *string             `json:"transport_address,omitempty"`
}

func (s *NodeTasks) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeTasks() *NodeTasks { _ = "STUB: not implemented"; return nil }
