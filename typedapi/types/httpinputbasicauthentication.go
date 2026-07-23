package types

type HttpInputBasicAuthentication struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func (s *HttpInputBasicAuthentication) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHttpInputBasicAuthentication() *HttpInputBasicAuthentication {
	_ = "STUB: not implemented"
	return nil
}

type HttpInputBasicAuthenticationVariant interface {
	HttpInputBasicAuthenticationCaster() *HttpInputBasicAuthentication
}

func (s *HttpInputBasicAuthentication) HttpInputBasicAuthenticationCaster() *HttpInputBasicAuthentication {
	_ = "STUB: not implemented"
	return nil
}
