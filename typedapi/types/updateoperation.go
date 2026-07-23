package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type UpdateOperation struct {
	Id_           *string `json:"_id,omitempty"`
	IfPrimaryTerm *int64  `json:"if_primary_term,omitempty"`
	IfSeqNo       *int64  `json:"if_seq_no,omitempty"`

	Index_ *string `json:"_index,omitempty"`

	RequireAlias *bool `json:"require_alias,omitempty"`

	RetryOnConflict *int `json:"retry_on_conflict,omitempty"`

	Routing     *string                  `json:"routing,omitempty"`
	Version     *int64                   `json:"version,omitempty"`
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *UpdateOperation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUpdateOperation() *UpdateOperation { _ = "STUB: not implemented"; return nil }

type UpdateOperationVariant interface {
	UpdateOperationCaster() *UpdateOperation
}

func (s *UpdateOperation) UpdateOperationCaster() *UpdateOperation {
	_ = "STUB: not implemented"
	return nil
}
