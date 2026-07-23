package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

type IndexAction struct {
	DocId              *string          `json:"doc_id,omitempty"`
	ExecutionTimeField *string          `json:"execution_time_field,omitempty"`
	Index              string           `json:"index"`
	OpType             *optype.OpType   `json:"op_type,omitempty"`
	Refresh            *refresh.Refresh `json:"refresh,omitempty"`
	Timeout            Duration         `json:"timeout,omitempty"`
}

func (s *IndexAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexAction() *IndexAction { _ = "STUB: not implemented"; return nil }

type IndexActionVariant interface {
	IndexActionCaster() *IndexAction
}

func (s *IndexAction) IndexActionCaster() *IndexAction { _ = "STUB: not implemented"; return nil }
