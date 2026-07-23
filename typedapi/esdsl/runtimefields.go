package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _runtimeFields struct {
	v types.RuntimeFields
}

func NewRuntimeFields(runtimefields map[string]types.RuntimeField) *_runtimeFields {
	_ = "STUB: not implemented"
	return nil
}

func (u *_runtimeFields) RuntimeFieldsCaster() *types.RuntimeFields {
	_ = "STUB: not implemented"
	return nil
}
