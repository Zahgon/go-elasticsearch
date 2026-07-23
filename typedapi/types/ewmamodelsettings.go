package types

type EwmaModelSettings struct {
	Alpha *float32 `json:"alpha,omitempty"`
}

func (s *EwmaModelSettings) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEwmaModelSettings() *EwmaModelSettings { _ = "STUB: not implemented"; return nil }

type EwmaModelSettingsVariant interface {
	EwmaModelSettingsCaster() *EwmaModelSettings
}

func (s *EwmaModelSettings) EwmaModelSettingsCaster() *EwmaModelSettings {
	_ = "STUB: not implemented"
	return nil
}
