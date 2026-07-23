package restore

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Accepted *bool                  `json:"accepted,omitempty"`
	Snapshot *types.SnapshotRestore `json:"snapshot,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
