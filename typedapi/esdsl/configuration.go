package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _configuration struct {
	v *types.Configuration
}

func NewConfiguration() *_configuration { _ = "STUB: not implemented"; return nil }

func (s *_configuration) FeatureStates(featurestates ...string) *_configuration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_configuration) IgnoreUnavailable(ignoreunavailable bool) *_configuration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_configuration) IncludeGlobalState(includeglobalstate bool) *_configuration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_configuration) Indices(indices ...string) *_configuration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_configuration) Metadata(metadata types.MetadataVariant) *_configuration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_configuration) Partial(partial bool) *_configuration {
	_ = "STUB: not implemented"
	return nil
}

func (s *_configuration) ConfigurationCaster() *types.Configuration {
	_ = "STUB: not implemented"
	return nil
}
