package types

type DiscoveryNodeCompact struct {
	Attributes       map[string]string `json:"attributes"`
	EphemeralId      string            `json:"ephemeral_id"`
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	TransportAddress string            `json:"transport_address"`
}

func (s *DiscoveryNodeCompact) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDiscoveryNodeCompact() *DiscoveryNodeCompact { _ = "STUB: not implemented"; return nil }
