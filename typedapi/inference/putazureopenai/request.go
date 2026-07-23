package putazureopenai

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/azureopenaiservicetype"
)

type Request struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service azureopenaiservicetype.AzureOpenAIServiceType `json:"service"`

	ServiceSettings types.AzureOpenAIServiceSettings `json:"service_settings"`

	TaskSettings *types.AzureOpenAITaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
