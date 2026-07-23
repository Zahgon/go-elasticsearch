package getinfluencers

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int64 `json:"count"`

	Influencers []types.Influencer `json:"influencers"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
