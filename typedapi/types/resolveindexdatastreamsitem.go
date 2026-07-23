package types

type ResolveIndexDataStreamsItem struct {
	BackingIndices []string `json:"backing_indices"`
	Name           string   `json:"name"`
	TimestampField string   `json:"timestamp_field"`
}

func (s *ResolveIndexDataStreamsItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewResolveIndexDataStreamsItem() *ResolveIndexDataStreamsItem {
	_ = "STUB: not implemented"
	return nil
}
