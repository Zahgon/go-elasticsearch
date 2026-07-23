package syncjobupdatestats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	DeletedDocumentCount int64 `json:"deleted_document_count"`

	IndexedDocumentCount int64 `json:"indexed_document_count"`

	IndexedDocumentVolume int64 `json:"indexed_document_volume"`

	LastSeen types.Duration `json:"last_seen,omitempty"`

	Metadata types.Metadata `json:"metadata,omitempty"`

	TotalDocumentCount *int `json:"total_document_count,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
