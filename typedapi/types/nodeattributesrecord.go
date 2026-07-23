package types

type NodeAttributesRecord struct {
	Attr *string `json:"attr,omitempty"`

	Host *string `json:"host,omitempty"`

	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Node *string `json:"node,omitempty"`

	Pid *string `json:"pid,omitempty"`

	Port *string `json:"port,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (s *NodeAttributesRecord) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeAttributesRecord() *NodeAttributesRecord { _ = "STUB: not implemented"; return nil }
