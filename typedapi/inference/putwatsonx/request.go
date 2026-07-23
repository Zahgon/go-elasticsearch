package putwatsonx

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/watsonxservicetype"
)

type Request struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service watsonxservicetype.WatsonxServiceType `json:"service"`

	ServiceSettings types.WatsonxServiceSettings `json:"service_settings"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
