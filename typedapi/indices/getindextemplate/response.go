package getindextemplate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	IndexTemplates []types.IndexTemplateItem `json:"index_templates"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
