package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _eSQLParams struct {
	v types.ESQLParams
}

func NewESQLParams() *_eSQLParams { _ = "STUB: not implemented"; return nil }

func (u *_eSQLParams) SingleOrMultiValues(singleormultivalues ...[]types.FieldValue) *_eSQLParams {
	_ = "STUB: not implemented"
	return nil
}

func (u *_eSQLParams) NamedValues(namedvalues ...types.NamedValueVariant) *_eSQLParams {
	_ = "STUB: not implemented"
	return nil
}

func (u *_eSQLParams) ESQLParamsCaster() *types.ESQLParams { _ = "STUB: not implemented"; return nil }
