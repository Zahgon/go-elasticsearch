package starttrainedmodeldeployment

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Assignment types.TrainedModelAssignment `json:"assignment"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
