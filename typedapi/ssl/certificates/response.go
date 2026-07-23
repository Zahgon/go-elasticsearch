package certificates

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response []types.CertificateInformation

func NewResponse() Response { _ = "STUB: not implemented"; return *new(Response) }
