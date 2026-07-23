package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensetype"
)

type License struct {
	ExpiryDateInMillis int64                   `json:"expiry_date_in_millis"`
	IssueDateInMillis  int64                   `json:"issue_date_in_millis"`
	IssuedTo           string                  `json:"issued_to"`
	Issuer             string                  `json:"issuer"`
	MaxNodes           *int64                  `json:"max_nodes,omitempty"`
	MaxResourceUnits   *int64                  `json:"max_resource_units,omitempty"`
	Signature          string                  `json:"signature"`
	StartDateInMillis  *int64                  `json:"start_date_in_millis,omitempty"`
	Type               licensetype.LicenseType `json:"type"`
	Uid                string                  `json:"uid"`
}

func (s *License) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLicense() *License { _ = "STUB: not implemented"; return nil }

type LicenseVariant interface {
	LicenseCaster() *License
}

func (s *License) LicenseCaster() *License { _ = "STUB: not implemented"; return nil }
