package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _definition struct {
	v *types.Definition
}

func NewDefinition(trainedmodel types.TrainedModelVariant) *_definition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_definition) Preprocessors(preprocessors ...types.PreprocessorVariant) *_definition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_definition) PreprocessorsValues(preprocessorsvalues []types.Preprocessor) *_definition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_definition) TrainedModel(trainedmodel types.TrainedModelVariant) *_definition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_definition) DefinitionCaster() *types.Definition { _ = "STUB: not implemented"; return nil }
