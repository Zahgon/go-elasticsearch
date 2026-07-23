package types

type DatabaseConfiguration struct {
	Ipinfo  *Ipinfo  `json:"ipinfo,omitempty"`
	Maxmind *Maxmind `json:"maxmind,omitempty"`

	Name string `json:"name"`
}

func (s *DatabaseConfiguration) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDatabaseConfiguration() *DatabaseConfiguration { _ = "STUB: not implemented"; return nil }

type DatabaseConfigurationVariant interface {
	DatabaseConfigurationCaster() *DatabaseConfiguration
}

func (s *DatabaseConfiguration) DatabaseConfigurationCaster() *DatabaseConfiguration {
	_ = "STUB: not implemented"
	return nil
}
