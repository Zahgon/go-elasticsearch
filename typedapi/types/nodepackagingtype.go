package types

type NodePackagingType struct {
	Count int `json:"count"`

	Flavor string `json:"flavor"`

	Type string `json:"type"`
}

func (s *NodePackagingType) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodePackagingType() *NodePackagingType { _ = "STUB: not implemented"; return nil }
