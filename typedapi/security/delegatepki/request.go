package delegatepki

type Request struct {
	X509CertificateChain []string `json:"x509_certificate_chain"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
