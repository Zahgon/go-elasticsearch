package putgroq

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/groqservicetype"
)

type Request struct {
	Service groqservicetype.GroqServiceType `json:"service"`

	ServiceSettings types.GroqServiceSettings `json:"service_settings"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
