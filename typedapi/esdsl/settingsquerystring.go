package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _settingsQueryString struct {
	v *types.SettingsQueryString
}

func NewSettingsQueryString() *_settingsQueryString { _ = "STUB: not implemented"; return nil }

func (s *_settingsQueryString) Lenient(stringifiedboolean types.StringifiedbooleanVariant) *_settingsQueryString {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsQueryString) SettingsQueryStringCaster() *types.SettingsQueryString {
	_ = "STUB: not implemented"
	return nil
}
