package putopenai

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openaiservicetype"
)

type Request struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service openaiservicetype.OpenAIServiceType `json:"service"`

	ServiceSettings types.OpenAIServiceSettings `json:"service_settings"`

	TaskSettings *types.OpenAITaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
