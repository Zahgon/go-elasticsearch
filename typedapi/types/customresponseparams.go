package types

type CustomResponseParams struct {
	JsonParser map[string]string `json:"json_parser"`
}

func NewCustomResponseParams() *CustomResponseParams { _ = "STUB: not implemented"; return nil }

type CustomResponseParamsVariant interface {
	CustomResponseParamsCaster() *CustomResponseParams
}

func (s *CustomResponseParams) CustomResponseParamsCaster() *CustomResponseParams {
	_ = "STUB: not implemented"
	return nil
}
