package grantapikey

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/apikeygranttype"
)

type Request struct {
	AccessToken *string `json:"access_token,omitempty"`

	ApiKey types.GrantApiKey `json:"api_key"`

	GrantType apikeygranttype.ApiKeyGrantType `json:"grant_type"`

	Password *string `json:"password,omitempty"`

	RunAs *string `json:"run_as,omitempty"`

	Username *string `json:"username,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
