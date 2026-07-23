package enrollnode

type Response struct {
	HttpCaCert string `json:"http_ca_cert"`

	HttpCaKey string `json:"http_ca_key"`

	NodesAddresses []string `json:"nodes_addresses"`

	TransportCaCert string `json:"transport_ca_cert"`

	TransportCert string `json:"transport_cert"`

	TransportKey string `json:"transport_key"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
