package types

type SecuritySettings struct {
	Index *IndexSettings `json:"index,omitempty"`
}

func NewSecuritySettings() *SecuritySettings { _ = "STUB: not implemented"; return nil }

type SecuritySettingsVariant interface {
	SecuritySettingsCaster() *SecuritySettings
}

func (s *SecuritySettings) SecuritySettingsCaster() *SecuritySettings {
	_ = "STUB: not implemented"
	return nil
}
