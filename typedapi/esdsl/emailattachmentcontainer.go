package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _emailAttachmentContainer struct {
	v *types.EmailAttachmentContainer
}

func NewEmailAttachmentContainer() *_emailAttachmentContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_emailAttachmentContainer) Data(data types.DataEmailAttachmentVariant) *_emailAttachmentContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_emailAttachmentContainer) Http(http types.HttpEmailAttachmentVariant) *_emailAttachmentContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_emailAttachmentContainer) Reporting(reporting types.ReportingEmailAttachmentVariant) *_emailAttachmentContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_emailAttachmentContainer) EmailAttachmentContainerCaster() *types.EmailAttachmentContainer {
	_ = "STUB: not implemented"
	return nil
}
