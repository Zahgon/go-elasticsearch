package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/managedby"
)

type DataStreamIndex struct {
	IlmPolicy *string `json:"ilm_policy,omitempty"`

	IndexMode *indexmode.IndexMode `json:"index_mode,omitempty"`

	IndexName string `json:"index_name"`

	IndexUuid string `json:"index_uuid"`

	ManagedBy *managedby.ManagedBy `json:"managed_by,omitempty"`

	PreferIlm *bool `json:"prefer_ilm,omitempty"`
}

func (s *DataStreamIndex) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDataStreamIndex() *DataStreamIndex { _ = "STUB: not implemented"; return nil }
