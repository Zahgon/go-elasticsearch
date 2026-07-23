package putazureaistudio

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/azureaistudioservicetype"
)

type Request struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service azureaistudioservicetype.AzureAiStudioServiceType `json:"service"`

	ServiceSettings types.AzureAiStudioServiceSettings `json:"service_settings"`

	TaskSettings *types.AzureAiStudioTaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
