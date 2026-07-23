package post

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensestatus"
)

type Response struct {
	Acknowledge   *types.Acknowledgement      `json:"acknowledge,omitempty"`
	Acknowledged  bool                        `json:"acknowledged"`
	LicenseStatus licensestatus.LicenseStatus `json:"license_status"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
