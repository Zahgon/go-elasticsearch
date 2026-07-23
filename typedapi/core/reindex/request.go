package reindex

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conflicts"
)

type Request struct {
	Conflicts *conflicts.Conflicts `json:"conflicts,omitempty"`

	Dest types.ReindexDestination `json:"dest"`

	MaxDocs *int64 `json:"max_docs,omitempty"`

	Script *types.Script `json:"script,omitempty"`

	Source types.ReindexSource `json:"source"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
