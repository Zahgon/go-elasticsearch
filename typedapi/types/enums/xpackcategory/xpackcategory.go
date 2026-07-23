package xpackcategory

type XPackCategory struct {
	Name string
}

var (
	Build = XPackCategory{"build"}

	Features = XPackCategory{"features"}

	License = XPackCategory{"license"}
)

func (x XPackCategory) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *XPackCategory) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (x XPackCategory) String() string { _ = "STUB: not implemented"; return "" }
