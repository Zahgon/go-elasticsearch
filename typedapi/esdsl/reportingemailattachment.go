package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _reportingEmailAttachment struct {
	v *types.ReportingEmailAttachment
}

func NewReportingEmailAttachment(url string) *_reportingEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reportingEmailAttachment) Inline(inline bool) *_reportingEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reportingEmailAttachment) Interval(duration types.DurationVariant) *_reportingEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reportingEmailAttachment) Request(request types.HttpInputRequestDefinitionVariant) *_reportingEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reportingEmailAttachment) Retries(retries int) *_reportingEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reportingEmailAttachment) Url(url string) *_reportingEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reportingEmailAttachment) EmailAttachmentContainerCaster() *types.EmailAttachmentContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reportingEmailAttachment) ReportingEmailAttachmentCaster() *types.ReportingEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}
