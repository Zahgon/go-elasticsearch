package types

type Monitoring struct {
	Available         bool             `json:"available"`
	CollectionEnabled bool             `json:"collection_enabled"`
	Enabled           bool             `json:"enabled"`
	EnabledExporters  map[string]int64 `json:"enabled_exporters"`
}

func (s *Monitoring) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMonitoring() *Monitoring { _ = "STUB: not implemented"; return nil }
