package putfireworksai

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fireworksaiservicetype"
)

type Request struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service fireworksaiservicetype.FireworksAIServiceType `json:"service"`

	ServiceSettings types.FireworksAIServiceSettings `json:"service_settings"`

	TaskSettings *types.FireworksAITaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
