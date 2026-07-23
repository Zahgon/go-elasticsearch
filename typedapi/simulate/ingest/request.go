package ingest

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	ComponentTemplateSubstitutions map[string]types.ComponentTemplateNode `json:"component_template_substitutions,omitempty"`

	Docs []types.Document `json:"docs"`

	IndexTemplateSubstitutions map[string]types.IndexTemplate `json:"index_template_substitutions,omitempty"`
	MappingAddition            *types.TypeMapping             `json:"mapping_addition,omitempty"`

	PipelineSubstitutions map[string]types.IngestPipeline `json:"pipeline_substitutions,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
