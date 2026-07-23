package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensestatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensetype"
)

type MinimalLicenseInformation struct {
	ExpiryDateInMillis int64                       `json:"expiry_date_in_millis"`
	Mode               licensetype.LicenseType     `json:"mode"`
	Status             licensestatus.LicenseStatus `json:"status"`
	Type               licensetype.LicenseType     `json:"type"`
	Uid                string                      `json:"uid"`
}

func (s *MinimalLicenseInformation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMinimalLicenseInformation() *MinimalLicenseInformation {
	_ = "STUB: not implemented"
	return nil
}
