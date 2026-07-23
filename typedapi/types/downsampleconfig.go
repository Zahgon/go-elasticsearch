package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/samplingmethod"
)

type DownsampleConfig struct {
	FixedInterval string `json:"fixed_interval"`

	SamplingMethod *samplingmethod.SamplingMethod `json:"sampling_method,omitempty"`
}

func (s *DownsampleConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDownsampleConfig() *DownsampleConfig { _ = "STUB: not implemented"; return nil }

type DownsampleConfigVariant interface {
	DownsampleConfigCaster() *DownsampleConfig
}

func (s *DownsampleConfig) DownsampleConfigCaster() *DownsampleConfig {
	_ = "STUB: not implemented"
	return nil
}
