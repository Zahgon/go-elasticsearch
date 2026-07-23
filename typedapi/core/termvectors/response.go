package termvectors

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Found       bool                        `json:"found"`
	Id_         *string                     `json:"_id,omitempty"`
	Index_      string                      `json:"_index"`
	TermVectors map[string]types.TermVector `json:"term_vectors,omitempty"`
	Took        int64                       `json:"took"`
	Version_    int64                       `json:"_version"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
