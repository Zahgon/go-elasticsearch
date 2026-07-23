package infertrainedmodel

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	InferenceResults []types.InferenceResponseResult `json:"inference_results"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
