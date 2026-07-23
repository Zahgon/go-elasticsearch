package types

type IndexState struct {
	Aliases    map[string]Alias `json:"aliases,omitempty"`
	DataStream *string          `json:"data_stream,omitempty"`

	Defaults *IndexSettings `json:"defaults,omitempty"`

	Lifecycle *DataStreamLifecycle `json:"lifecycle,omitempty"`
	Mappings  *TypeMapping         `json:"mappings,omitempty"`
	Settings  *IndexSettings       `json:"settings,omitempty"`
}

func (s *IndexState) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexState() *IndexState { _ = "STUB: not implemented"; return nil }
