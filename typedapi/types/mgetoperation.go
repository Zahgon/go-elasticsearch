package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type MgetOperation struct {
	Id_ string `json:"_id"`

	Index_ *string `json:"_index,omitempty"`

	Routing []string `json:"routing,omitempty"`

	Source_ SourceConfig `json:"_source,omitempty"`

	StoredFields []string                 `json:"stored_fields,omitempty"`
	Version      *int64                   `json:"version,omitempty"`
	VersionType  *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *MgetOperation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMgetOperation() *MgetOperation { _ = "STUB: not implemented"; return nil }

type MgetOperationVariant interface {
	MgetOperationCaster() *MgetOperation
}

func (s *MgetOperation) MgetOperationCaster() *MgetOperation { _ = "STUB: not implemented"; return nil }
