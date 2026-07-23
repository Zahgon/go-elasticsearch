package putnvidia

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/nvidiaservicetype"
)

type Request struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service nvidiaservicetype.NvidiaServiceType `json:"service"`

	ServiceSettings types.NvidiaServiceSettings `json:"service_settings"`

	TaskSettings *types.NvidiaTaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
