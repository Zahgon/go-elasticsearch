package activateuserprofile

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/granttype"
)

type Request struct {
	AccessToken *string `json:"access_token,omitempty"`

	GrantType granttype.GrantType `json:"grant_type"`

	Password *string `json:"password,omitempty"`

	Username *string `json:"username,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
