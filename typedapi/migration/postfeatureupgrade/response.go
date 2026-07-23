package postfeatureupgrade

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Accepted bool                         `json:"accepted"`
	Features []types.PostMigrationFeature `json:"features,omitempty"`
	Reason   *string                      `json:"reason,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
