package imageurldetail

type ImageUrlDetail struct {
	Name string
}

var (
	Auto = ImageUrlDetail{"auto"}

	Low = ImageUrlDetail{"low"}

	High = ImageUrlDetail{"high"}
)

func (i ImageUrlDetail) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *ImageUrlDetail) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i ImageUrlDetail) String() string { _ = "STUB: not implemented"; return "" }
