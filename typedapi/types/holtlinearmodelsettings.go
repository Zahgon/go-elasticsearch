package types

type HoltLinearModelSettings struct {
	Alpha *float32 `json:"alpha,omitempty"`
	Beta  *float32 `json:"beta,omitempty"`
}

func (s *HoltLinearModelSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHoltLinearModelSettings() *HoltLinearModelSettings { _ = "STUB: not implemented"; return nil }

type HoltLinearModelSettingsVariant interface {
	HoltLinearModelSettingsCaster() *HoltLinearModelSettings
}

func (s *HoltLinearModelSettings) HoltLinearModelSettingsCaster() *HoltLinearModelSettings {
	_ = "STUB: not implemented"
	return nil
}
