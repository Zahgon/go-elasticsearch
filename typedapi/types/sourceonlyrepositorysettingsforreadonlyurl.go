package types

type SourceOnlyRepositorySettingsForReadOnlyUrl struct {
	DelegateType string `json:"delegate_type,omitempty"`

	HttpMaxRetries *int `json:"http_max_retries,omitempty"`

	HttpSocketTimeout Duration `json:"http_socket_timeout,omitempty"`

	MaxNumberOfSnapshots *int `json:"max_number_of_snapshots,omitempty"`

	Url string `json:"url"`
}

func (s *SourceOnlyRepositorySettingsForReadOnlyUrl) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SourceOnlyRepositorySettingsForReadOnlyUrl) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSourceOnlyRepositorySettingsForReadOnlyUrl() *SourceOnlyRepositorySettingsForReadOnlyUrl {
	_ = "STUB: not implemented"
	return nil
}

type SourceOnlyRepositorySettingsForReadOnlyUrlVariant interface {
	SourceOnlyRepositorySettingsForReadOnlyUrlCaster() *SourceOnlyRepositorySettingsForReadOnlyUrl
}

func (s *SourceOnlyRepositorySettingsForReadOnlyUrl) SourceOnlyRepositorySettingsForReadOnlyUrlCaster() *SourceOnlyRepositorySettingsForReadOnlyUrl {
	_ = "STUB: not implemented"
	return nil
}

func (s *SourceOnlyRepositorySettingsForReadOnlyUrl) SourceOnlyRepositorySettingsCaster() *SourceOnlyRepositorySettings {
	_ = "STUB: not implemented"
	return nil
}
