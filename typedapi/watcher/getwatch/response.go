package getwatch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Found        bool               `json:"found"`
	Id_          string             `json:"_id"`
	PrimaryTerm_ *int               `json:"_primary_term,omitempty"`
	SeqNo_       *int64             `json:"_seq_no,omitempty"`
	Status       *types.WatchStatus `json:"status,omitempty"`
	Version_     *int64             `json:"_version,omitempty"`
	Watch        *types.Watch       `json:"watch,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
