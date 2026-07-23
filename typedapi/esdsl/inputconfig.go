package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _inputConfig struct {
	v *types.InputConfig
}

func NewInputConfig(inputfield string, outputfield string) *_inputConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inputConfig) InputField(inputfield string) *_inputConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inputConfig) OutputField(outputfield string) *_inputConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inputConfig) InputConfigCaster() *types.InputConfig {
	_ = "STUB: not implemented"
	return nil
}
