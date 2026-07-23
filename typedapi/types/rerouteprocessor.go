package types

type RerouteProcessor struct {
	Dataset []string `json:"dataset,omitempty"`

	Description *string `json:"description,omitempty"`

	Destination *string `json:"destination,omitempty"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	Namespace []string `json:"namespace,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`
}

func (s *RerouteProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRerouteProcessor() *RerouteProcessor { _ = "STUB: not implemented"; return nil }

type RerouteProcessorVariant interface {
	RerouteProcessorCaster() *RerouteProcessor
}

func (s *RerouteProcessor) RerouteProcessorCaster() *RerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}
