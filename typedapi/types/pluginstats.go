package types

type PluginStats struct {
	Classname            string   `json:"classname"`
	Description          string   `json:"description"`
	ElasticsearchVersion string   `json:"elasticsearch_version"`
	ExtendedPlugins      []string `json:"extended_plugins"`
	HasNativeController  bool     `json:"has_native_controller"`
	JavaVersion          string   `json:"java_version"`
	Licensed             bool     `json:"licensed"`
	Name                 string   `json:"name"`
	Version              string   `json:"version"`
}

func (s *PluginStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPluginStats() *PluginStats { _ = "STUB: not implemented"; return nil }
