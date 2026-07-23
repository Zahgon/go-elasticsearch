package putamazonsagemaker

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/amazonsagemakerservicetype"
)

type Request struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service amazonsagemakerservicetype.AmazonSageMakerServiceType `json:"service"`

	ServiceSettings types.AmazonSageMakerServiceSettings `json:"service_settings"`

	TaskSettings *types.AmazonSageMakerTaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
