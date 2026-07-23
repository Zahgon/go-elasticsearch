package types

type FielddataRecord struct {
	Field *string `json:"field,omitempty"`

	Host *string `json:"host,omitempty"`

	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Node *string `json:"node,omitempty"`

	Size *string `json:"size,omitempty"`
}

func (s *FielddataRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFielddataRecord() *FielddataRecord { _ = "STUB: not implemented"; return nil }
