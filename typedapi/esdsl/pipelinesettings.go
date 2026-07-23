package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pipelineSettings struct {
	v *types.PipelineSettings
}

func NewPipelineSettings(pipelinebatchdelay int, pipelinebatchsize int, pipelineworkers int, queuecheckpointwrites int, queuemaxbytes string, queuetype string) *_pipelineSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineSettings) PipelineBatchDelay(pipelinebatchdelay int) *_pipelineSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineSettings) PipelineBatchSize(pipelinebatchsize int) *_pipelineSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineSettings) PipelineWorkers(pipelineworkers int) *_pipelineSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineSettings) QueueCheckpointWrites(queuecheckpointwrites int) *_pipelineSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineSettings) QueueMaxBytes(queuemaxbytes string) *_pipelineSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineSettings) QueueType(queuetype string) *_pipelineSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineSettings) PipelineSettingsCaster() *types.PipelineSettings {
	_ = "STUB: not implemented"
	return nil
}
