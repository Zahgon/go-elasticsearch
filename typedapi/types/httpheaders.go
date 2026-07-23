package types

type HttpHeaders map[string][]string

type HttpHeadersVariant interface {
	HttpHeadersCaster() *HttpHeaders
}
