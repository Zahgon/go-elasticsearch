package types

type FilteringConfig struct {
	Active FilteringRules `json:"active"`
	Domain *string        `json:"domain,omitempty"`
	Draft  FilteringRules `json:"draft"`
}

func (s *FilteringConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFilteringConfig() *FilteringConfig { _ = "STUB: not implemented"; return nil }

type FilteringConfigVariant interface {
	FilteringConfigCaster() *FilteringConfig
}

func (s *FilteringConfig) FilteringConfigCaster() *FilteringConfig {
	_ = "STUB: not implemented"
	return nil
}
