package types

type NetworkDirectionProcessor struct {
	Description *string `json:"description,omitempty"`

	DestinationIp *string `json:"destination_ip,omitempty"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	InternalNetworks []string `json:"internal_networks,omitempty"`

	InternalNetworksField *string `json:"internal_networks_field,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	SourceIp *string `json:"source_ip,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *NetworkDirectionProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNetworkDirectionProcessor() *NetworkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

type NetworkDirectionProcessorVariant interface {
	NetworkDirectionProcessorCaster() *NetworkDirectionProcessor
}

func (s *NetworkDirectionProcessor) NetworkDirectionProcessorCaster() *NetworkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}
