package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pipelineMetadata struct {
	v *types.PipelineMetadata
}

func NewPipelineMetadata(type_ string, version string) *_pipelineMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineMetadata) Type(type_ string) *_pipelineMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineMetadata) Version(version string) *_pipelineMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineMetadata) PipelineMetadataCaster() *types.PipelineMetadata {
	_ = "STUB: not implemented"
	return nil
}
