package putgooglevertexai

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/googlevertexaiservicetype"
)

type Request struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service googlevertexaiservicetype.GoogleVertexAIServiceType `json:"service"`

	ServiceSettings types.GoogleVertexAIServiceSettings `json:"service_settings"`

	TaskSettings *types.GoogleVertexAITaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
