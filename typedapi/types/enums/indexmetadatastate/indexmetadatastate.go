package indexmetadatastate

type IndexMetadataState struct {
	Name string
}

var (
	Open = IndexMetadataState{"open"}

	Close = IndexMetadataState{"close"}
)

func (i IndexMetadataState) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IndexMetadataState) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (i IndexMetadataState) String() string { _ = "STUB: not implemented"; return "" }
