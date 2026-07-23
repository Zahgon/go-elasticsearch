package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type MTermVectorsOperation struct {
	Doc json.RawMessage `json:"doc,omitempty"`

	FieldStatistics *bool `json:"field_statistics,omitempty"`

	Fields []string `json:"fields,omitempty"`

	Filter *TermVectorsFilter `json:"filter,omitempty"`

	Id_ *string `json:"_id,omitempty"`

	Index_ *string `json:"_index,omitempty"`

	Offsets *bool `json:"offsets,omitempty"`

	Payloads *bool `json:"payloads,omitempty"`

	Positions *bool `json:"positions,omitempty"`

	Routing []string `json:"routing,omitempty"`

	TermStatistics *bool `json:"term_statistics,omitempty"`

	Version *int64 `json:"version,omitempty"`

	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *MTermVectorsOperation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMTermVectorsOperation() *MTermVectorsOperation { _ = "STUB: not implemented"; return nil }

type MTermVectorsOperationVariant interface {
	MTermVectorsOperationCaster() *MTermVectorsOperation
}

func (s *MTermVectorsOperation) MTermVectorsOperationCaster() *MTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}
