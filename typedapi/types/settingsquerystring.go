package types

type SettingsQueryString struct {
	Lenient Stringifiedboolean `json:"lenient"`
}

func (s *SettingsQueryString) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSettingsQueryString() *SettingsQueryString { _ = "STUB: not implemented"; return nil }

type SettingsQueryStringVariant interface {
	SettingsQueryStringCaster() *SettingsQueryString
}

func (s *SettingsQueryString) SettingsQueryStringCaster() *SettingsQueryString {
	_ = "STUB: not implemented"
	return nil
}
