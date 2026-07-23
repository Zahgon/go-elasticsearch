package updatesettings

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Security *types.SecuritySettings `json:"security,omitempty"`

	SecurityProfile *types.SecuritySettings `json:"security-profile,omitempty"`

	SecurityTokens *types.SecuritySettings `json:"security-tokens,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
