package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _elasticsearchServiceSettings struct {
	v *types.ElasticsearchServiceSettings
}

func NewElasticsearchServiceSettings(modelid string, numthreads int) *_elasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elasticsearchServiceSettings) AdaptiveAllocations(adaptiveallocations types.AdaptiveAllocationsVariant) *_elasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elasticsearchServiceSettings) DeploymentId(deploymentid string) *_elasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elasticsearchServiceSettings) LongDocumentStrategy(longdocumentstrategy string) *_elasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elasticsearchServiceSettings) MaxChunksPerDoc(maxchunksperdoc int) *_elasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elasticsearchServiceSettings) ModelId(modelid string) *_elasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elasticsearchServiceSettings) NumAllocations(numallocations int) *_elasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elasticsearchServiceSettings) NumThreads(numthreads int) *_elasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elasticsearchServiceSettings) ElasticsearchServiceSettingsCaster() *types.ElasticsearchServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
