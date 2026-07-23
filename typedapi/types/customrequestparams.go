package types

type CustomRequestParams struct {
	Content string `json:"content"`
}

func (s *CustomRequestParams) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCustomRequestParams() *CustomRequestParams { _ = "STUB: not implemented"; return nil }

type CustomRequestParamsVariant interface {
	CustomRequestParamsCaster() *CustomRequestParams
}

func (s *CustomRequestParams) CustomRequestParamsCaster() *CustomRequestParams {
	_ = "STUB: not implemented"
	return nil
}
