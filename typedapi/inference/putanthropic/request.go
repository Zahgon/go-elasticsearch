package putanthropic

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/anthropicservicetype"
)

type Request struct {
	Service anthropicservicetype.AnthropicServiceType `json:"service"`

	ServiceSettings types.AnthropicServiceSettings `json:"service_settings"`

	TaskSettings *types.AnthropicTaskSettings `json:"task_settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
