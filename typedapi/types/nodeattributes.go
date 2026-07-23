package types

type NodeAttributes struct {
	Attributes map[string]string `json:"attributes"`

	EphemeralId string `json:"ephemeral_id"`

	Id *string `json:"id,omitempty"`

	Name string `json:"name"`

	TransportAddress string `json:"transport_address"`
}

func (s *NodeAttributes) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeAttributes() *NodeAttributes { _ = "STUB: not implemented"; return nil }
