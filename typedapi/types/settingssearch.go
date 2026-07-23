package types

type SettingsSearch struct {
	Idle    *SearchIdle      `json:"idle,omitempty"`
	Slowlog *SlowlogSettings `json:"slowlog,omitempty"`
}

func NewSettingsSearch() *SettingsSearch { _ = "STUB: not implemented"; return nil }

type SettingsSearchVariant interface {
	SettingsSearchCaster() *SettingsSearch
}

func (s *SettingsSearch) SettingsSearchCaster() *SettingsSearch {
	_ = "STUB: not implemented"
	return nil
}
