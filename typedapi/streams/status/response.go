package status

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Logs     types.StreamStatus `json:"logs"`
	LogsEcs  types.StreamStatus `json:"logs.ecs"`
	LogsOtel types.StreamStatus `json:"logs.otel"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
