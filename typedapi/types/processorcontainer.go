package types

import (
	"encoding/json"
)

type ProcessorContainer struct {
	AdditionalProcessorContainerProperty map[string]json.RawMessage `json:"-"`

	Append *AppendProcessor `json:"append,omitempty"`

	Attachment *AttachmentProcessor `json:"attachment,omitempty"`

	Bytes *BytesProcessor `json:"bytes,omitempty"`

	Cef *CefProcessor `json:"cef,omitempty"`

	Circle *CircleProcessor `json:"circle,omitempty"`

	CommunityId *CommunityIDProcessor `json:"community_id,omitempty"`

	Convert *ConvertProcessor `json:"convert,omitempty"`

	Csv *CsvProcessor `json:"csv,omitempty"`

	Date *DateProcessor `json:"date,omitempty"`

	DateIndexName *DateIndexNameProcessor `json:"date_index_name,omitempty"`

	Dissect *DissectProcessor `json:"dissect,omitempty"`

	DotExpander *DotExpanderProcessor `json:"dot_expander,omitempty"`

	Drop *DropProcessor `json:"drop,omitempty"`

	Enrich *EnrichProcessor `json:"enrich,omitempty"`

	Fail *FailProcessor `json:"fail,omitempty"`

	Fingerprint *FingerprintProcessor `json:"fingerprint,omitempty"`

	Foreach *ForeachProcessor `json:"foreach,omitempty"`

	GeoGrid *GeoGridProcessor `json:"geo_grid,omitempty"`

	Geoip *GeoIpProcessor `json:"geoip,omitempty"`

	Grok *GrokProcessor `json:"grok,omitempty"`

	Gsub *GsubProcessor `json:"gsub,omitempty"`

	HtmlStrip *HtmlStripProcessor `json:"html_strip,omitempty"`

	Inference *InferenceProcessor `json:"inference,omitempty"`

	IpLocation *IpLocationProcessor `json:"ip_location,omitempty"`

	Join *JoinProcessor `json:"join,omitempty"`

	Json *JsonProcessor `json:"json,omitempty"`

	Kv *KeyValueProcessor `json:"kv,omitempty"`

	Lowercase *LowercaseProcessor `json:"lowercase,omitempty"`

	NetworkDirection *NetworkDirectionProcessor `json:"network_direction,omitempty"`

	Pipeline *PipelineProcessor `json:"pipeline,omitempty"`

	Redact *RedactProcessor `json:"redact,omitempty"`

	RegisteredDomain *RegisteredDomainProcessor `json:"registered_domain,omitempty"`

	Remove *RemoveProcessor `json:"remove,omitempty"`

	Rename *RenameProcessor `json:"rename,omitempty"`

	Reroute *RerouteProcessor `json:"reroute,omitempty"`

	Script *ScriptProcessor `json:"script,omitempty"`

	Set *SetProcessor `json:"set,omitempty"`

	SetSecurityUser *SetSecurityUserProcessor `json:"set_security_user,omitempty"`

	Sort *SortProcessor `json:"sort,omitempty"`

	Split *SplitProcessor `json:"split,omitempty"`

	Terminate *TerminateProcessor `json:"terminate,omitempty"`

	Trim *TrimProcessor `json:"trim,omitempty"`

	Uppercase *UppercaseProcessor `json:"uppercase,omitempty"`

	UriParts *UriPartsProcessor `json:"uri_parts,omitempty"`

	Urldecode *UrlDecodeProcessor `json:"urldecode,omitempty"`

	UserAgent *UserAgentProcessor `json:"user_agent,omitempty"`
}

func (s *ProcessorContainer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ProcessorContainer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewProcessorContainer() *ProcessorContainer { _ = "STUB: not implemented"; return nil }

type ProcessorContainerVariant interface {
	ProcessorContainerCaster() *ProcessorContainer
}

func (s *ProcessorContainer) ProcessorContainerCaster() *ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}
