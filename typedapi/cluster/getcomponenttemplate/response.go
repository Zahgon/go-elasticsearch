package getcomponenttemplate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ComponentTemplates []types.ClusterComponentTemplate `json:"component_templates"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
