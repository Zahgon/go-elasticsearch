package httpinputmethod

type HttpInputMethod struct {
	Name string
}

var (
	Head = HttpInputMethod{"head"}

	Get = HttpInputMethod{"get"}

	Post = HttpInputMethod{"post"}

	Put = HttpInputMethod{"put"}

	Delete = HttpInputMethod{"delete"}
)

func (h HttpInputMethod) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HttpInputMethod) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (h HttpInputMethod) String() string { _ = "STUB: not implemented"; return "" }
