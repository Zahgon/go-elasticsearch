package dataattachmentformat

type DataAttachmentFormat struct {
	Name string
}

var (
	Json = DataAttachmentFormat{"json"}

	Yaml = DataAttachmentFormat{"yaml"}
)

func (d DataAttachmentFormat) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DataAttachmentFormat) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DataAttachmentFormat) String() string { _ = "STUB: not implemented"; return "" }
