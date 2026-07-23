package gettrainedmodelsstats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int `json:"count"`

	TrainedModelStats []types.TrainedModelStats `json:"trained_model_stats"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
