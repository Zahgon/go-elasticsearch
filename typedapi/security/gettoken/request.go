package gettoken

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/accesstokengranttype"
)

type Request struct {
	GrantType *accesstokengranttype.AccessTokenGrantType `json:"grant_type,omitempty"`

	KerberosTicket *string `json:"kerberos_ticket,omitempty"`

	Password *string `json:"password,omitempty"`

	RefreshToken *string `json:"refresh_token,omitempty"`

	Scope *string `json:"scope,omitempty"`

	Username *string `json:"username,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
