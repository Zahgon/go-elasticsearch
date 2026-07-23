package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _settingsSearch struct {
	v *types.SettingsSearch
}

func NewSettingsSearch() *_settingsSearch { _ = "STUB: not implemented"; return nil }

func (s *_settingsSearch) Idle(idle types.SearchIdleVariant) *_settingsSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSearch) Slowlog(slowlog types.SlowlogSettingsVariant) *_settingsSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSearch) SettingsSearchCaster() *types.SettingsSearch {
	_ = "STUB: not implemented"
	return nil
}
