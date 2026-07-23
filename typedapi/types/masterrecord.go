package types

type MasterRecord struct {
	Host *string `json:"host,omitempty"`

	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Node *string `json:"node,omitempty"`
}

func (s *MasterRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMasterRecord() *MasterRecord { _ = "STUB: not implemented"; return nil }
