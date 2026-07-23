package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _chainInput struct {
	v *types.ChainInput
}

func NewChainInput() *_chainInput { _ = "STUB: not implemented"; return nil }

func (s *_chainInput) Inputs(inputs []map[string]types.WatcherInput) *_chainInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chainInput) WatcherInputCaster() *types.WatcherInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chainInput) ChainInputCaster() *types.ChainInput { _ = "STUB: not implemented"; return nil }
