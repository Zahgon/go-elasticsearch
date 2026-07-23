package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _customResponseParams struct {
	v *types.CustomResponseParams
}

func NewCustomResponseParams() *_customResponseParams { _ = "STUB: not implemented"; return nil }

func (s *_customResponseParams) JsonParser(jsonparser map[string]string) *_customResponseParams {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customResponseParams) AddJsonParser(key string, value string) *_customResponseParams {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customResponseParams) CustomResponseParamsCaster() *types.CustomResponseParams {
	_ = "STUB: not implemented"
	return nil
}
