package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noderole"
)

type CurrentNode struct {
	Attributes       map[string]string   `json:"attributes"`
	Id               string              `json:"id"`
	Name             string              `json:"name"`
	Roles            []noderole.NodeRole `json:"roles"`
	TransportAddress string              `json:"transport_address"`
	WeightRanking    int                 `json:"weight_ranking"`
}

func (s *CurrentNode) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCurrentNode() *CurrentNode { _ = "STUB: not implemented"; return nil }
