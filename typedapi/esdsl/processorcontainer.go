package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _processorContainer struct {
	v *types.ProcessorContainer
}

func NewProcessorContainer() *_processorContainer { _ = "STUB: not implemented"; return nil }

func (s *_processorContainer) AdditionalProcessorContainerProperty(key string, value json.RawMessage) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Append(append types.AppendProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Attachment(attachment types.AttachmentProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Bytes(bytes types.BytesProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Cef(cef types.CefProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Circle(circle types.CircleProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) CommunityId(communityid types.CommunityIDProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Convert(convert types.ConvertProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Csv(csv types.CsvProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Date(date types.DateProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) DateIndexName(dateindexname types.DateIndexNameProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Dissect(dissect types.DissectProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) DotExpander(dotexpander types.DotExpanderProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Drop(drop types.DropProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Enrich(enrich types.EnrichProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Fail(fail types.FailProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Fingerprint(fingerprint types.FingerprintProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Foreach(foreach types.ForeachProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) GeoGrid(geogrid types.GeoGridProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Geoip(geoip types.GeoIpProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Grok(grok types.GrokProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Gsub(gsub types.GsubProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) HtmlStrip(htmlstrip types.HtmlStripProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Inference(inference types.InferenceProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) IpLocation(iplocation types.IpLocationProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Join(join types.JoinProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Json(json types.JsonProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Kv(kv types.KeyValueProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Lowercase(lowercase types.LowercaseProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) NetworkDirection(networkdirection types.NetworkDirectionProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Pipeline(pipeline types.PipelineProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Redact(redact types.RedactProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) RegisteredDomain(registereddomain types.RegisteredDomainProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Remove(remove types.RemoveProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Rename(rename types.RenameProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Reroute(reroute types.RerouteProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Script(script types.ScriptProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Set(set types.SetProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) SetSecurityUser(setsecurityuser types.SetSecurityUserProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Sort(sort types.SortProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Split(split types.SplitProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Terminate(terminate types.TerminateProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Trim(trim types.TrimProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Uppercase(uppercase types.UppercaseProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) UriParts(uriparts types.UriPartsProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) Urldecode(urldecode types.UrlDecodeProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) UserAgent(useragent types.UserAgentProcessorVariant) *_processorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_processorContainer) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}
