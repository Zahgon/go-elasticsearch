package types

type DatabaseConfigurationFull struct {
	Ipinfo  *Ipinfo  `json:"ipinfo,omitempty"`
	Local   *Local   `json:"local,omitempty"`
	Maxmind *Maxmind `json:"maxmind,omitempty"`

	Name string `json:"name"`
	Web  *Web   `json:"web,omitempty"`
}

func (s *DatabaseConfigurationFull) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDatabaseConfigurationFull() *DatabaseConfigurationFull {
	_ = "STUB: not implemented"
	return nil
}
