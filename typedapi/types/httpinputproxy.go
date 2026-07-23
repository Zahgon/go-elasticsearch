package types

type HttpInputProxy struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func (s *HttpInputProxy) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHttpInputProxy() *HttpInputProxy { _ = "STUB: not implemented"; return nil }

type HttpInputProxyVariant interface {
	HttpInputProxyCaster() *HttpInputProxy
}

func (s *HttpInputProxy) HttpInputProxyCaster() *HttpInputProxy {
	_ = "STUB: not implemented"
	return nil
}
