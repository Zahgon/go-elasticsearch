package putopenshiftai

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openshiftaiservicetype"
)

type Request struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	Service openshiftaiservicetype.OpenShiftAiServiceType `json:"service"`

	ServiceSettings types.OpenShiftAiServiceSettings `json:"service_settings"`

	TaskSettings *types.OpenShiftAiTaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
