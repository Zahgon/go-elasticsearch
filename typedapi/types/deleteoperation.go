package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type DeleteOperation struct {
	Id_           *string `json:"_id,omitempty"`
	IfPrimaryTerm *int64  `json:"if_primary_term,omitempty"`
	IfSeqNo       *int64  `json:"if_seq_no,omitempty"`

	Index_ *string `json:"_index,omitempty"`

	Routing     *string                  `json:"routing,omitempty"`
	Version     *int64                   `json:"version,omitempty"`
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *DeleteOperation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDeleteOperation() *DeleteOperation { _ = "STUB: not implemented"; return nil }

type DeleteOperationVariant interface {
	DeleteOperationCaster() *DeleteOperation
}

func (s *DeleteOperation) DeleteOperationCaster() *DeleteOperation {
	_ = "STUB: not implemented"
	return nil
}
