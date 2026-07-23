package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _securitySettings struct {
	v *types.SecuritySettings
}

func NewSecuritySettings() *_securitySettings { _ = "STUB: not implemented"; return nil }

func (s *_securitySettings) Index(index types.IndexSettingsVariant) *_securitySettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_securitySettings) SecuritySettingsCaster() *types.SecuritySettings {
	_ = "STUB: not implemented"
	return nil
}
