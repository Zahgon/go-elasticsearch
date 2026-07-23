package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ingestPipelineParams struct {
	v *types.IngestPipelineParams
}

func NewIngestPipelineParams(extractbinarycontent bool, name string, reducewhitespace bool, runmlinference bool) *_ingestPipelineParams {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipelineParams) ExtractBinaryContent(extractbinarycontent bool) *_ingestPipelineParams {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipelineParams) Name(name string) *_ingestPipelineParams {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipelineParams) ReduceWhitespace(reducewhitespace bool) *_ingestPipelineParams {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipelineParams) RunMlInference(runmlinference bool) *_ingestPipelineParams {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipelineParams) IngestPipelineParamsCaster() *types.IngestPipelineParams {
	_ = "STUB: not implemented"
	return nil
}
