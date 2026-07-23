package suggestuserprofiles

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Profiles []types.UserProfile `json:"profiles"`

	Took int64 `json:"took"`

	Total types.TotalUserProfiles `json:"total"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
