package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noderole"
)

type SearchShardsNodeAttributes struct {
	Attributes map[string]string `json:"attributes"`

	EphemeralId     string `json:"ephemeral_id"`
	ExternalId      string `json:"external_id"`
	MaxIndexVersion int    `json:"max_index_version"`
	MinIndexVersion int    `json:"min_index_version"`

	Name  string              `json:"name"`
	Roles []noderole.NodeRole `json:"roles"`

	TransportAddress string `json:"transport_address"`
	Version          string `json:"version"`
}

func (s *SearchShardsNodeAttributes) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSearchShardsNodeAttributes() *SearchShardsNodeAttributes {
	_ = "STUB: not implemented"
	return nil
}
