package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noderole"
)

type UpdateByQueryRethrottleNode struct {
	Attributes       map[string]string   `json:"attributes"`
	Host             string              `json:"host"`
	Ip               string              `json:"ip"`
	Name             string              `json:"name"`
	Roles            []noderole.NodeRole `json:"roles,omitempty"`
	Tasks            map[string]TaskInfo `json:"tasks"`
	TransportAddress string              `json:"transport_address"`
}

func (s *UpdateByQueryRethrottleNode) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUpdateByQueryRethrottleNode() *UpdateByQueryRethrottleNode {
	_ = "STUB: not implemented"
	return nil
}
