package responsecontenttype

type ResponseContentType struct {
	Name string
}

var (
	Json = ResponseContentType{"json"}

	Yaml = ResponseContentType{"yaml"}

	Text = ResponseContentType{"text"}
)

func (r ResponseContentType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ResponseContentType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (r ResponseContentType) String() string { _ = "STUB: not implemented"; return "" }
