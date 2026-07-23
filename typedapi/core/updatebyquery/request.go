package updatebyquery

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conflicts"
)

type Request struct {
	Conflicts *conflicts.Conflicts `json:"conflicts,omitempty"`

	MaxDocs *int64 `json:"max_docs,omitempty"`

	Query *types.Query `json:"query,omitempty"`

	Script *types.Script `json:"script,omitempty"`

	Slice *types.SlicedScroll `json:"slice,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
