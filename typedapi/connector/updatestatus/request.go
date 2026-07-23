package updatestatus

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectorstatus"
)

type Request struct {
	Status connectorstatus.ConnectorStatus `json:"status"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
