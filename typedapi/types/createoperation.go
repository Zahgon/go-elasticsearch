package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type CreateOperation struct {
	DynamicTemplates map[string]string `json:"dynamic_templates,omitempty"`

	Id_           *string `json:"_id,omitempty"`
	IfPrimaryTerm *int64  `json:"if_primary_term,omitempty"`
	IfSeqNo       *int64  `json:"if_seq_no,omitempty"`

	Index_ *string `json:"_index,omitempty"`

	Pipeline *string `json:"pipeline,omitempty"`

	RequireAlias *bool `json:"require_alias,omitempty"`

	Routing     *string                  `json:"routing,omitempty"`
	Version     *int64                   `json:"version,omitempty"`
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *CreateOperation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCreateOperation() *CreateOperation { _ = "STUB: not implemented"; return nil }

type CreateOperationVariant interface {
	CreateOperationCaster() *CreateOperation
}

func (s *CreateOperation) CreateOperationCaster() *CreateOperation {
	_ = "STUB: not implemented"
	return nil
}
