package putcontextualai

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/contextualaiservicetype"
)

type Request struct {
	Service contextualaiservicetype.ContextualAIServiceType `json:"service"`

	ServiceSettings types.ContextualAIServiceSettings `json:"service_settings"`

	TaskSettings *types.ContextualAITaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
