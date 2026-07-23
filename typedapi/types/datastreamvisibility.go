package types

type DataStreamVisibility struct {
	AllowCustomRouting *bool `json:"allow_custom_routing,omitempty"`
	FailureStore       *bool `json:"failure_store,omitempty"`
	Hidden             *bool `json:"hidden,omitempty"`
}

func (s *DataStreamVisibility) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamVisibility() *DataStreamVisibility { _ = "STUB: not implemented"; return nil }

type DataStreamVisibilityVariant interface {
	DataStreamVisibilityCaster() *DataStreamVisibility
}

func (s *DataStreamVisibility) DataStreamVisibilityCaster() *DataStreamVisibility {
	_ = "STUB: not implemented"
	return nil
}
