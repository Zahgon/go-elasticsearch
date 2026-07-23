package gettrainedmodels

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int `json:"count"`

	TrainedModelConfigs []types.TrainedModelConfig `json:"trained_model_configs"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
