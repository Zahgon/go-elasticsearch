package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _datafeedConfig struct {
	v *types.DatafeedConfig
}

func NewDatafeedConfig() *_datafeedConfig { _ = "STUB: not implemented"; return nil }

func (s *_datafeedConfig) Aggregations(aggregations map[string]types.Aggregations) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) AddAggregation(key string, value types.AggregationsVariant) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) ChunkingConfig(chunkingconfig types.ChunkingConfigVariant) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) DatafeedId(id string) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) DelayedDataCheckConfig(delayeddatacheckconfig types.DelayedDataCheckConfigVariant) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) Frequency(duration types.DurationVariant) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) Indices(indices ...string) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) IndicesOptions(indicesoptions types.IndicesOptionsVariant) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) JobId(id string) *_datafeedConfig { _ = "STUB: not implemented"; return nil }

func (s *_datafeedConfig) MaxEmptySearches(maxemptysearches int) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) Query(query types.QueryVariant) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) QueryDelay(duration types.DurationVariant) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) ScriptFields(scriptfields map[string]types.ScriptField) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) AddScriptField(key string, value types.ScriptFieldVariant) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) ScrollSize(scrollsize int) *_datafeedConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_datafeedConfig) DatafeedConfigCaster() *types.DatafeedConfig {
	_ = "STUB: not implemented"
	return nil
}
