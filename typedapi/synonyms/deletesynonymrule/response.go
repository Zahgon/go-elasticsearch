package deletesynonymrule

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/result"
)

type Response struct {
	ReloadAnalyzersDetails *types.ReloadResult `json:"reload_analyzers_details,omitempty"`

	Result result.Result `json:"result"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
