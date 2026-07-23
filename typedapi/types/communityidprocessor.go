package types

type CommunityIDProcessor struct {
	Description *string `json:"description,omitempty"`

	DestinationIp *string `json:"destination_ip,omitempty"`

	DestinationPort *string `json:"destination_port,omitempty"`

	IanaNumber *string `json:"iana_number,omitempty"`

	IcmpCode *string `json:"icmp_code,omitempty"`

	IcmpType *string `json:"icmp_type,omitempty"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Seed *int `json:"seed,omitempty"`

	SourceIp *string `json:"source_ip,omitempty"`

	SourcePort *string `json:"source_port,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`

	Transport *string `json:"transport,omitempty"`
}

func (s *CommunityIDProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCommunityIDProcessor() *CommunityIDProcessor { _ = "STUB: not implemented"; return nil }

type CommunityIDProcessorVariant interface {
	CommunityIDProcessorCaster() *CommunityIDProcessor
}

func (s *CommunityIDProcessor) CommunityIDProcessorCaster() *CommunityIDProcessor {
	_ = "STUB: not implemented"
	return nil
}
