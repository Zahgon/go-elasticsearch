package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensestatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensetype"
)

type LicenseInformation struct {
	ExpiryDate DateTime `json:"expiry_date,omitempty"`

	ExpiryDateInMillis *int64 `json:"expiry_date_in_millis,omitempty"`

	IssueDate DateTime `json:"issue_date"`

	IssueDateInMillis int64 `json:"issue_date_in_millis"`

	IssuedTo string `json:"issued_to"`

	Issuer string `json:"issuer"`

	MaxNodes *int64 `json:"max_nodes,omitempty"`

	MaxResourceUnits *int `json:"max_resource_units,omitempty"`

	StartDateInMillis int64 `json:"start_date_in_millis"`

	Status licensestatus.LicenseStatus `json:"status"`

	Type licensetype.LicenseType `json:"type"`

	Uid string `json:"uid"`
}

func (s *LicenseInformation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewLicenseInformation() *LicenseInformation { _ = "STUB: not implemented"; return nil }
