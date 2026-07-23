package types

type RRFRetrieverComponent struct {
	Retriever RetrieverContainer `json:"retriever"`

	Weight *float32 `json:"weight,omitempty"`
}

func (s *RRFRetrieverComponent) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRRFRetrieverComponent() *RRFRetrieverComponent { _ = "STUB: not implemented"; return nil }

type RRFRetrieverComponentVariant interface {
	RRFRetrieverComponentCaster() *RRFRetrieverComponent
}

func (s *RRFRetrieverComponent) RRFRetrieverComponentCaster() *RRFRetrieverComponent {
	_ = "STUB: not implemented"
	return nil
}

func (s *RRFRetrieverComponent) RRFRetrieverEntryCaster() *RRFRetrieverEntry {
	_ = "STUB: not implemented"
	return nil
}
