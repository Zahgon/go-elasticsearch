package activateuserprofile

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Data             map[string]json.RawMessage   `json:"data"`
	Doc_             types.UserProfileHitMetadata `json:"_doc"`
	Enabled          *bool                        `json:"enabled,omitempty"`
	Labels           map[string]json.RawMessage   `json:"labels"`
	LastSynchronized int64                        `json:"last_synchronized"`
	Uid              string                       `json:"uid"`
	User             types.UserProfileUser        `json:"user"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
