package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensetype"
)

type _license struct {
	v *types.License
}

func NewLicense(issuedto string, issuer string, signature string, type_ licensetype.LicenseType, uid string) *_license {
	_ = "STUB: not implemented"
	return nil
}

func (s *_license) ExpiryDateInMillis(epochtimeunitmillis int64) *_license {
	_ = "STUB: not implemented"
	return nil
}

func (s *_license) IssueDateInMillis(epochtimeunitmillis int64) *_license {
	_ = "STUB: not implemented"
	return nil
}

func (s *_license) IssuedTo(issuedto string) *_license { _ = "STUB: not implemented"; return nil }

func (s *_license) Issuer(issuer string) *_license { _ = "STUB: not implemented"; return nil }

func (s *_license) MaxNodes(maxnodes int64) *_license { _ = "STUB: not implemented"; return nil }

func (s *_license) MaxResourceUnits(maxresourceunits int64) *_license {
	_ = "STUB: not implemented"
	return nil
}

func (s *_license) Signature(signature string) *_license { _ = "STUB: not implemented"; return nil }

func (s *_license) StartDateInMillis(epochtimeunitmillis int64) *_license {
	_ = "STUB: not implemented"
	return nil
}

func (s *_license) Type(type_ licensetype.LicenseType) *_license {
	_ = "STUB: not implemented"
	return nil
}

func (s *_license) Uid(uid string) *_license { _ = "STUB: not implemented"; return nil }

func (s *_license) LicenseCaster() *types.License { _ = "STUB: not implemented"; return nil }
