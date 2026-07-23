package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type ReindexDestination struct {
	Index string `json:"index"`

	OpType *optype.OpType `json:"op_type,omitempty"`

	Pipeline *string `json:"pipeline,omitempty"`

	Routing *string `json:"routing,omitempty"`

	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *ReindexDestination) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewReindexDestination() *ReindexDestination { _ = "STUB: not implemented"; return nil }

type ReindexDestinationVariant interface {
	ReindexDestinationCaster() *ReindexDestination
}

func (s *ReindexDestination) ReindexDestinationCaster() *ReindexDestination {
	_ = "STUB: not implemented"
	return nil
}
