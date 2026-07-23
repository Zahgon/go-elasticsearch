package restmethod

type RestMethod struct {
	Name string
}

var (
	GET = RestMethod{"GET"}

	HEAD = RestMethod{"HEAD"}

	POST = RestMethod{"POST"}

	PUT = RestMethod{"PUT"}

	DELETE = RestMethod{"DELETE"}
)

func (r RestMethod) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RestMethod) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (r RestMethod) String() string { _ = "STUB: not implemented"; return "" }
