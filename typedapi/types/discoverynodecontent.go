package types

type DiscoveryNodeContent struct {
	Attributes       map[string]string `json:"attributes"`
	EphemeralId      string            `json:"ephemeral_id"`
	ExternalId       string            `json:"external_id"`
	MaxIndexVersion  int               `json:"max_index_version"`
	MinIndexVersion  int               `json:"min_index_version"`
	Name             *string           `json:"name,omitempty"`
	Roles            []string          `json:"roles"`
	TransportAddress string            `json:"transport_address"`
	Version          string            `json:"version"`
}

func (s *DiscoveryNodeContent) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDiscoveryNodeContent() *DiscoveryNodeContent { _ = "STUB: not implemented"; return nil }
