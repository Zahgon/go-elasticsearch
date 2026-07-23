package types

type HttpInputAuthentication struct {
	Basic HttpInputBasicAuthentication `json:"basic"`
}

func NewHttpInputAuthentication() *HttpInputAuthentication { _ = "STUB: not implemented"; return nil }

type HttpInputAuthenticationVariant interface {
	HttpInputAuthenticationCaster() *HttpInputAuthentication
}

func (s *HttpInputAuthentication) HttpInputAuthenticationCaster() *HttpInputAuthentication {
	_ = "STUB: not implemented"
	return nil
}
