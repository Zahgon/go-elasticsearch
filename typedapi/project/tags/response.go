package tags

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	LinkedProjects map[string]types.Tags `json:"linked_projects,omitempty"`
	Origin         map[string]types.Tags `json:"origin"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
