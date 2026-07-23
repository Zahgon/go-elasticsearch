package contenttype

type ContentType struct {
	Name string
}

var (
	Text = ContentType{"text"}

	Imageurl = ContentType{"image_url"}

	File = ContentType{"file"}
)

func (c ContentType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ContentType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ContentType) String() string { _ = "STUB: not implemented"; return "" }
