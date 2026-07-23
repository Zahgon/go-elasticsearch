package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sourceConfig struct {
	v types.SourceConfig
}

func NewSourceConfig() *_sourceConfig { _ = "STUB: not implemented"; return nil }

func (u *_sourceConfig) Bool(bool bool) *_sourceConfig { _ = "STUB: not implemented"; return nil }

func (u *_sourceConfig) SourceFilter(sourcefilter types.SourceFilterVariant) *_sourceConfig {
	_ = "STUB: not implemented"
	return nil
}

func (u *_sourceFilter) SourceConfigCaster() *types.SourceConfig {
	_ = "STUB: not implemented"
	return nil
}

func (u *_sourceConfig) SourceConfigCaster() *types.SourceConfig {
	_ = "STUB: not implemented"
	return nil
}
