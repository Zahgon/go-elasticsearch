package putdeepseek

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/deepseekservicetype"
)

type Request struct {
	Service deepseekservicetype.DeepSeekServiceType `json:"service"`

	ServiceSettings types.DeepSeekServiceSettings `json:"service_settings"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
