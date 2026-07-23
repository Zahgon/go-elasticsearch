package mltrainedmodels

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response []types.TrainedModelsRecord

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
