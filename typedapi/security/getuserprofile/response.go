package getuserprofile

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Errors *types.GetUserProfileErrors `json:"errors,omitempty"`

	Profiles []types.UserProfileWithMetadata `json:"profiles"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
