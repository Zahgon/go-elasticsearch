package putcohere

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cohereservicetype"
)

type Request struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service cohereservicetype.CohereServiceType `json:"service"`

	ServiceSettings types.CohereServiceSettings `json:"service_settings"`

	TaskSettings *types.CohereTaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
